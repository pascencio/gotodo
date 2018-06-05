package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// EchoServer server of the app
type EchoServer struct {
	ResourceDefinitions []ResourceDefinition
}

// EntityValidationError validation error on request entity
type EntityValidationError struct {
	err     string
	details map[string]string
}

func (e EntityValidationError) Error() string {
	return e.err
}

// Details get entity validation detail
func (e EntityValidationError) Details() map[string]string {
	return e.details
}

// EchoRequestHandler request handler for REST services
type EchoRequestHandler struct {
	Context  echo.Context
	Resource Resource
}

// NewRequestContext factory for RequestContext
func (h EchoRequestHandler) NewRequestContext() RequestContext {
	requestContext := RequestContext{
		pathParamHandler: func(name string) string {
			return h.Context.Param(name)
		},
		queryParamHandler: func(name string) string {
			return h.Context.QueryParam(name)
		},
		entityHandler: func(entity interface{}) error {
			defer h.Context.Request().Body.Close()
			bytes, err := ioutil.ReadAll(h.Context.Request().Body)

			if err != nil {
				return err
			}

			json.Unmarshal(bytes, &entity)

			_, err = govalidator.ValidateStruct(entity)

			if err != nil {
				details := map[string]string{}
				message := "Validations error on resource"
				switch errType := err.(type) {
				case govalidator.Error:
					valErr := err.(govalidator.Error)
					details[valErr.Name] = formatValidationError(valErr)
				case govalidator.Errors:
					for _, e := range errType {
						valErr := e.(govalidator.Error)
						details[valErr.Name] = formatValidationError(valErr)
					}
				}

				return EntityValidationError{
					err:     message,
					details: details,
				}

			}

			log.WithField("entity", entity).Debug("Get entity complete")

			return nil
		},
	}
	return requestContext
}

func formatValidationError(valError govalidator.Error) string {
	pattern := viper.GetString("messages.validation." + valError.Validator)
	if pattern != "" {
		return fmt.Sprintf(pattern, valError.Name)
	}
	return valError.Error()
}

// Run start the echo server
func (s EchoServer) Run() {
	e := echo.New()

	for _, definition := range s.ResourceDefinitions {
		for _, resource := range definition.Resources {
			switch resource.GetMethod() {
			case http.MethodGet:
				e.GET(buildResourcePath(definition, resource), createRequestHandler(resource))
			case http.MethodPut:
				e.PUT(buildResourcePath(definition, resource), createRequestHandler(resource))
			case http.MethodPost:
				e.POST(buildResourcePath(definition, resource), createRequestHandler(resource))
			case http.MethodDelete:
				e.DELETE(buildResourcePath(definition, resource), createRequestHandler(resource))
			}
		}
	}

	e.Start(viper.GetString("address"))
}

func createRequestHandler(resource Resource) func(c echo.Context) error {
	return func(c echo.Context) error {
		echoHandler := EchoRequestHandler{
			Context:  c,
			Resource: resource,
		}
		requestContext := echoHandler.NewRequestContext()

		log.WithFields(log.Fields{
			"method": resource.GetMethod(),
			"path":   resource.GetPath(),
		}).Debug("Handling request")

		output, err := resource.Handler(requestContext)

		if err != nil {
			output := map[string]interface{}{}

			if valErr, ok := err.(ValidationError); ok {
				output["errors"] = valErr.Details()
			} else {
				output["errors"] = map[string]string{
					"internal": err.Error(),
				}
			}

			return c.JSON(http.StatusBadRequest, output)
		}

		if r := recover(); r != nil {
			log.Error(r)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if output == nil {
			output = make(map[string]string)
		}

		if k := reflect.TypeOf(output).Kind(); k == reflect.Slice || k == reflect.Array {
			if v := reflect.ValueOf(output); v.Len() == 0 {
				output = []map[string]string{}
			}
		}

		return c.JSON(http.StatusOK, output)
	}

}

func buildResourcePath(definition ResourceDefinition, resource Resource) string {
	if resource.GetPath() != "" {
		return fmt.Sprintf("/%s/%s", definition.Path, resource.GetPath())
	}
	return fmt.Sprintf("/%s", definition.Path)
}
