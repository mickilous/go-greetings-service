package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	serverPort int
	logger     *log.Logger
	router     *mux.Router
}

func NewServer(serverPort int, logger *log.Logger, router *mux.Router) *Server {
	return &Server{
		serverPort: serverPort,
		logger:     logger,
		router:     router,
	}
}

func (s *Server) Start() {
	s.logger.Printf("Starting server on %v", s.serverPort)
	s.logger.Fatal(
		"Startup of server failed",
		http.ListenAndServe(":"+strconv.Itoa(s.serverPort), s.router))
}
