package http

import (
	"github.com/gorilla/mux"
	"greetings-service/configuration"
	"log"
	"net/http"
)

type Server struct {
	*log.Logger
	*mux.Router
	ServerPort string
}

func NewServer(configurationProvider configuration.Provider, logger *log.Logger, router *mux.Router) *Server {
	return &Server{
		logger,
		router,
		configurationProvider.GetString("SERVER_PORT", "8080")}
}

func (s *Server) Start() {
	s.Logger.Printf("Starting server on %v", s.ServerPort)
	s.Logger.Fatal(
		"Startup of server failed",
		http.ListenAndServe(":"+s.ServerPort, s.Router))
}
