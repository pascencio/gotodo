package rest

import (
	"net/http"

	"github.com/pascencio/gotodo/todo"
	log "github.com/sirupsen/logrus"
)

type HateoasSpec struct {
	Domain   interface{}
	LinkSpec []LinkSpec `json:"_link"`
}

type LinkSpec struct {
	Rel  string
	Href string
}

type RequestHandlerError struct {
	Err error
}

func NewCrudResource(path string, unmarshal func([]byte) interface{}) ResourceDefinition {

	resources := []Resource{
		Resource{
			Method:    http.MethodGet,
			Unmarshal: unmarshal,
			Handler: func(c RequestContext) interface{} {
				todo := todo.Todo{
					Title:       "Call to mom",
					Description: "Call to mom today",
				}

				todo.Id = "1"
				return todo
			},
		},
		Resource{
			Method:    http.MethodGet,
			Unmarshal: unmarshal,
			Path:      ":id",
			Handler: func(c RequestContext) interface{} {
				id := c.PathParam("id")
				log.WithFields(log.Fields{
					"id": id,
				}).Debug("Obteniendo todo por id")
				todo := todo.Todo{
					Title:       "Call to mom",
					Description: "Call to mom today",
				}
				todo.Id = id
				return todo
			},
		},
		Resource{
			Method: http.MethodPost,
			Handler: func(c RequestContext) interface{} {
				todo := &todo.Todo{}
				c.Entity(todo)

				log.WithFields(log.Fields{
					"todo": todo,
				}).Debug("Insertando todo")

				return todo
			},
		},
		Resource{
			Method: http.MethodPut,
			Path:   ":id",
			Handler: func(c RequestContext) interface{} {
				todo := &todo.Todo{}
				c.Entity(todo)
				todo.Id = c.PathParam("id")
				log.WithFields(log.Fields{
					"todo": todo,
				}).Debug("Actualizando todo")

				return todo
			},
		},
		Resource{
			Method: http.MethodDelete,
			Path:   ":id",
			Handler: func(c RequestContext) interface{} {
				todo := &todo.Todo{}
				c.Entity(todo)
				todo.Id = c.PathParam("id")
				log.WithFields(log.Fields{
					"todo": todo,
				}).Debug("Eliminando todo")

				return todo
			},
		},
	}

	definition := ResourceDefinition{
		Path:      path,
		Resources: resources,
	}

	return definition
}
