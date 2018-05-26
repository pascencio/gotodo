package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pascencio/gotodo/config"
	"github.com/pascencio/gotodo/repository"
	"github.com/pascencio/gotodo/rest"
	"github.com/pascencio/gotodo/todo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.gotodo")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error al leer archivo de configuración: %s", err))
	}
	level, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		panic(fmt.Errorf("Error al cargar nivel de logs: %s", err))
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}

// TodoApplication Configuration for startup and dependecy injection
type TodoApplication struct {
}

// ConnectionPool configuration for TodoApplication
func (c TodoApplication) ConnectionPool(context config.ConfigurationContext) repository.ConnectionPool {
	return repository.MongoConnectionPool{}
}

// Server configuraion for TodoApplication
func (c TodoApplication) Server(context config.ConfigurationContext) rest.Server {
	return rest.EchoServer{
		ResourceDefinitions: []rest.ResourceDefinition{
			rest.NewCrudResourceDefinition("todo", func(bytes []byte) interface{} {
				todo := &todo.Todo{}
				json.Unmarshal(bytes, &todo)
				return todo
			}),
		},
	}
}

// Start start TodoApplication
func (c TodoApplication) Start() {

	context := config.ConfigurationContext{}
	server := c.Server(context)
	server.Run()
}
