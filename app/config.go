package app

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo/bson"

	"github.com/pascencio/gotodo/config"
	"github.com/pascencio/gotodo/data"
	"github.com/pascencio/gotodo/data/mongodb"
	"github.com/pascencio/gotodo/deliver/http"
	"github.com/pascencio/gotodo/todo"
	"github.com/pascencio/gotodo/todo/mongo"
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
func (c TodoApplication) ConnectionPool(context config.ConfigurationContext) interface{} {
	return mongodb.MongoConnectionPool{}
}

// Server configuraion for TodoApplication
func (c TodoApplication) Server(context config.ConfigurationContext) interface{} {
	connection := context.BeanDefinitions["ConnectionPool"].GetBean(context).(mongodb.MongoConnectionPool)
	err := connection.Start()

	if err != nil {
		panic(fmt.Errorf("Error starting application: %s", err))
	}

	mongoTemplate := mongodb.MongoRepositoryTemplate{}

	mongoTemplate.SetConnection(&connection)

	todoRepository := mongo.TodoRepository{}
	todoRepository.Template = mongoTemplate
	return http.EchoServer{
		ResourceDefinitions: []http.ResourceDefinition{
			http.NewCrudResourceDefinition(
				"todo",
				todoRepository,
				func(r http.RequestContext) (data.Domain, error) {
					result := &todo.Todo{}

					err := r.Entity(result)

					return result, err
				},
				func(i data.Iterator) data.Domain {
					result := &todo.Todo{}
					if !i.Next(result) {
						return nil
					}
					return result
				},
				func(id string) interface{} {
					return bson.ObjectIdHex(id)
				},
			),
		},
	}
}

// Start start TodoApplication
func (c TodoApplication) Start() {

	context := config.ConfigurationContext{}
	context.BeanDefinitions = map[string]config.BeanDefinition{
		"ConnectionPool": config.BeanDefinition{
			Name:    "ConnectionPool",
			Scope:   config.ScopeSingleton,
			Factory: c.ConnectionPool,
		},
	}
	server := c.Server(context).(http.Server)
	server.Run()
}
