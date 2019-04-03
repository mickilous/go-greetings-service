package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var serverPort = getEnv("SERVER_PORT", "8080")

func Start(router *mux.Router, logger *log.Logger) {
	logger.Printf("Starting server on %v", serverPort)
	err := http.ListenAndServe(":"+serverPort, router)
	if err != nil {
		logger.Fatal("Startup of server failed", err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
