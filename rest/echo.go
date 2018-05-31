package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// EchoServer server of the app
type EchoServer struct {
	ResourceDefinitions []ResourceDefinition
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
		entityHandler: func(entity interface{}) {
			defer h.Context.Request().Body.Close()
			bytes, error := ioutil.ReadAll(h.Context.Request().Body)

			if error != nil {
				panic("Error al leer bytes de la entidad")
			}

			json.Unmarshal(bytes, &entity)

			log.WithField("entity", entity).Debug("Get entity complete")
		},
	}
	return requestContext
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

		output := resource.Handler(requestContext)

		if r := recover(); r != nil {
			log.Error(r)
			return echo.NewHTTPError(http.StatusInternalServerError)
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
