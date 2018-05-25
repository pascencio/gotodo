package rest

import (
	"encoding/json"

	"github.com/pascencio/gotodo/todo"
)

type Resource struct {
	Method    string
	Path      string
	Unmarshal func([]byte) interface{}
	Handler   func(RequestContext) interface{}
}

type ResourceDefinition struct {
	Path      string
	Resources []Resource
}

type ResourceConfigurator interface {
	Configure(ResourceDefinition)
}

func Setup(configurator ResourceConfigurator) {
	definitions := []ResourceDefinition{
		NewCrudResource("todo", func(bytes []byte) interface{} {
			todo := &todo.Todo{}
			json.Unmarshal(bytes, &todo)
			return todo
		}),
	}
	for _, definition := range definitions {
		configurator.Configure(definition)
	}
}
