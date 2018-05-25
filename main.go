package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/pascencio/gotodo/rest"
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
		panic(fmt.Errorf("Error al leer archivo de configuración: %s \n", err))
	}
	level, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		panic(fmt.Errorf("Error al cargar nivel de logs: %s \n", err))
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}

func main() {
	e := echo.New()
	configurator := rest.EchoResourceConfigurator{
		Echo: e,
	}
	rest.Setup(configurator)
	e.Start(viper.GetString("address"))
}
