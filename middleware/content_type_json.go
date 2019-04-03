package middleware

import (
	"net/http"
)

type ContentTypeJson struct{}

func NewContentTypeJson() *ContentTypeJson {
	return &ContentTypeJson{}
}

func (c *ContentTypeJson) HandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(writer, request)
	}
}
