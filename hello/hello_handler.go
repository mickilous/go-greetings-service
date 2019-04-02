package hello

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	logger *log.Logger
}

type Message struct {
	Message string `json:message`
}

var buddies = map[string]string{
	"042": "Marvin",
	"666": "Zuul",
	"007": "Bond",
}

func NewHelloHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/{version}/hello/{userId}", h.MiddleWare(h.Hello()))
}

func (h *Handler) Hello() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		version, userId := vars["version"], vars["userId"]
		var httpStatus int
		var message string
		switch version {
		case "v1":
			httpStatus = http.StatusOK
			message = fmt.Sprintf("Yo %v!", buddies[userId])
		case "v2":
			httpStatus = http.StatusOK
			message = fmt.Sprintf("Hello %v!", buddies[userId])
		case "v3":
			httpStatus = http.StatusOK
			message = fmt.Sprintf("How do you do %v!", buddies[userId])
		default:
			h.logger.Printf("Unsupported Version %v", version)
			httpStatus = http.StatusBadRequest
			message = fmt.Sprintf("Unsupported Version %v", version)
		}
		writer.WriteHeader(httpStatus)
		json.NewEncoder(writer).Encode(Message{Message: message})
		//writer.Write([]byte(message))
	}
}

func (h *Handler) MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		route, vars := mux.CurrentRoute(request), mux.Vars(request)
		pathRegexp, _ := route.GetPathTemplate()
		h.logger.Printf("Handling request to %v with parameters :", pathRegexp)
		for k, v := range vars {
			h.logger.Printf("\tkey: %v, value: %v", k, v)
		}
		before := time.Now()
		defer h.logger.Printf("Request processed in %v", time.Now().Sub(before))
		next(writer, request)
	}
}
