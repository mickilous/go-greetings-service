package main

import (
	"github.com/gorilla/mux"
	"greetings-service/configuration"
	"greetings-service/deserve"
	"greetings-service/greetings"
	"greetings-service/http"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "greet-srv ", log.LstdFlags|log.Lshortfile)

	configurationProvider := configuration.NewEnvironmentProvider()

	httpClient := http.NewClient()

	router := mux.NewRouter()

	deserveClient := deserve.NewClient(configurationProvider, logger, httpClient)

	greetingsHandler := greetings.NewHandlers(logger, deserveClient)
	greetingsHandler.SetupRoutes(router)

	http.NewServer(configurationProvider, logger, router).Start()

}
