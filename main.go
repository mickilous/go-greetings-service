package main

import (
	"github.com/gorilla/mux"
	"gitlab.com/mickilous/go-greetings-service/greetings"
	"gitlab.com/mickilous/go-greetings-service/server"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "greet-srv ", log.LstdFlags|log.Lshortfile)

	router := mux.NewRouter()

	helloHandler := greetings.NewHandler(logger)
	helloHandler.SetupRoutes(router)

	server.Start(router, logger)

}
