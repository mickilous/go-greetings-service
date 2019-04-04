package http

import (
	"net/http"
	"time"
)

func NewClient() *http.Client {
	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	return &httpClient
}
