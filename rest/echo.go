package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type EchoResourceConfigurator struct {
	Echo *echo.Echo
}

func (c EchoResourceConfigurator) Configure(definition ResourceDefinition) {
	for _, resource := range definition.Resources {
		switch resource.Method {
		case http.MethodGet:
			c.Echo.GET(buildResourcePath(definition, resource), createRequestHandler(resource))
		case http.MethodPut:
			c.Echo.PUT(buildResourcePath(definition, resource), createRequestHandler(resource))
		case http.MethodPost:
			c.Echo.POST(buildResourcePath(definition, resource), createRequestHandler(resource))
		case http.MethodDelete:
			c.Echo.DELETE(buildResourcePath(definition, resource), createRequestHandler(resource))
		}
	}
}

func createRequestHandler(resource Resource) func(c echo.Context) error {
	return func(c echo.Context) error {
		echoHandler := EchoRequestHandler{
			Context:  c,
			Resource: resource,
		}
		requestContext := echoHandler.NewRequestContext()

		output := resource.Handler(requestContext)

		if r := recover(); r != nil {
			log.Error(r)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, output)
	}

}

func buildResourcePath(definition ResourceDefinition, resource Resource) string {
	if resource.Path != "" {
		return fmt.Sprintf("/%s/%s", definition.Path, resource.Path)
	} else {
		return fmt.Sprintf("/%s", definition.Path)
	}
}

type EchoRequestHandler struct {
	Context  echo.Context
	Resource Resource
}

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

			log.WithFields(log.Fields{
				"entity": entity,
			}).Debug("Entidad obtenida")
		},
	}
	return requestContext
}
