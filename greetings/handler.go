package greetings

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"greetings-service/deserve"
	"greetings-service/middleware"
	"log"
	"net/http"
)

type Handlers struct {
	logger        *log.Logger
	deserveClient *deserve.Client
}

type Message struct {
	Message string `json:"message"`
}

var buddies = map[string]string{
	"042": "Marvin",
	"666": "Zuul",
	"007": "Bond",
}

func NewHandlers(logger *log.Logger, client *deserve.Client) *Handlers {
	return &Handlers{
		logger:        logger,
		deserveClient: client,
	}
}

func (h *Handlers) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/{version}/hello/{userId}", middleWares(h, h.Hello()))
}

func middleWares(h *Handlers, handlerFunc func(writer http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return middleware.NewContentTypeJson().HandlerFunc(middleware.NewRequestLogger(h.logger).HandlerFunc(handlerFunc))
}

func (h *Handlers) Hello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		version, userId := vars["version"], vars["userId"]
		isGreetable := h.deserveClient.IsGreetable(userId)
		var httpStatus int
		var ret string
		switch version {
		case "v1":
			httpStatus = http.StatusOK
			ret = buildMessage("Yo %v", userId, isGreetable)
		case "v2":
			httpStatus = http.StatusOK
			ret = buildMessage("Hello %v!", userId, isGreetable)
		case "v3":
			httpStatus = http.StatusOK
			ret = buildMessage("How do you do %v!", userId, isGreetable)
		default:
			http.Error(writer, fmt.Sprintf("Unsupported Version %v", version), http.StatusBadRequest)
			return
		}
		writer.WriteHeader(httpStatus)
		json.NewEncoder(writer).Encode(Message{Message: ret})
	}
}

func buildMessage(message string, userId string, isGreetable bool) string {
	if isGreetable {
		return fmt.Sprintf(message, buddies[userId])
	} else {
		return fmt.Sprintf("Go to hell %v!", buddies[userId])
	}
}
