package greetings

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

type message struct {
	message string `json:message`
}

var buddies = map[string]string{
	"042": "Marvin",
	"666": "Zuul",
	"007": "Bond",
}

func NewHandler(logger *log.Logger) *Handler {
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
		var ret string
		switch version {
		case "v1":
			httpStatus = http.StatusOK
			ret = fmt.Sprintf("Yo %v!", buddies[userId])
		case "v2":
			httpStatus = http.StatusOK
			ret = fmt.Sprintf("Hello %v!", buddies[userId])
		case "v3":
			httpStatus = http.StatusOK
			ret = fmt.Sprintf("How do you do %v!", buddies[userId])
		default:
			h.logger.Printf("Unsupported Version %v", version)
			httpStatus = http.StatusBadRequest
			ret = fmt.Sprintf("Unsupported Version %v", version)
		}
		writer.WriteHeader(httpStatus)
		json.NewEncoder(writer).Encode(message{message: ret})
	}
}

func (h *Handler) MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		route, vars := mux.CurrentRoute(request), mux.Vars(request)
		path, _ := route.GetPathTemplate()
		h.logger.Printf("Handling request to %v with parameters :", path)
		for key, val := range vars {
			h.logger.Printf("\t%v : %v", key, val)
		}
		before := time.Now()
		defer h.logger.Printf("Request processed in %v", time.Now().Sub(before))
		next(writer, request)
	}
}
