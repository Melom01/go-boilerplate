package main

import (
	"log"

	"github.com/Melom01/go-boilerplate/pkg/config"
	"github.com/Melom01/go-boilerplate/pkg/di"
)

func main() {
	configuration, configurationError := config.LoadConfiguration()
	if configurationError != nil {
		log.Fatal("Unable to load the server configuration: ", configurationError)
	}

	server, diError := di.InitializeAppDependencies(configuration)
	if diError != nil {
		log.Fatal("Unable to start the server: ", diError)
	} else {
		server.Start(configuration.ServerPort)
	}
}
