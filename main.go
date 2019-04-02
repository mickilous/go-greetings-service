package main

import (
	"github.com/gorilla/mux"
	"gitlab.com/mickilous/go-greetings-service/greetings"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "greet-srv ", log.LstdFlags|log.Lshortfile)

	server := mux.NewRouter()

	helloHandler := greetings.NewGreetingsHandler(logger)
	helloHandler.SetupRoutes(server)

	logger.Println("Starting server on 8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		logger.Fatal("Startup of server failed", err)
	}

}
