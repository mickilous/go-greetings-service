package middleware

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type RequestLogger struct {
	logger *log.Logger
}

func NewRequestLogger(logger *log.Logger) *RequestLogger {
	return &RequestLogger{logger: logger}
}

func (r *RequestLogger) HandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		route, vars := mux.CurrentRoute(request), mux.Vars(request)
		path, _ := route.GetPathTemplate()
		r.logger.Printf("Handling request to %v with parameters :", path)
		for key, val := range vars {
			r.logger.Printf("\t%v : %v", key, val)
		}
		before := time.Now()
		defer r.logger.Printf("Request processed in %v", time.Now().Sub(before))
		next(writer, request)
	}
}
