package main

import (
	"github.com/gorilla/mux"
	"greetings-service/deserve"
	"greetings-service/greetings"
	"greetings-service/http"
	"log"
	"os"
	"strconv"
)

var (
	serverPort  = getEnvInt("SERVER_PORT", 8080)
	deserveAddr = getEnvStr("DESERVE_ADDR", "http://localhost:8090")
	logger      *log.Logger
)

func main() {

	logger := log.New(os.Stdout, "greet-srv ", log.LstdFlags|log.Lshortfile)
	httpClient := http.NewClient()
	router := mux.NewRouter()
	deserveClient := deserve.NewClient(logger, httpClient, deserveAddr)
	greetingsHandler := greetings.NewHandler(logger, deserveClient)
	greetingsHandler.SetupRoutes(router)

	http.NewServer(serverPort, logger, router).Start()

}

func getEnvStr(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	ret, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		logger.Fatalf("Unsupported Int ENV Var - %v : %v", key, value)
	}
	return int(ret)

}
