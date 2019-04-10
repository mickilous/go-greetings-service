package greetings_test

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"greetings-service/configuration"
	"greetings-service/deserve"
	"greetings-service/greetings"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

type TestHelloData struct {
	version string
	userId  string
	deserve bool
	message string
}

var testHelloDatas = []TestHelloData{
	{"v1", "007", true, "Yo Bond!"},
	{"v2", "007", true, "Hello Bond!"},
	{"v3", "007", true, "How do you do Bond!"},
	{"v1", "042", true, "Yo Marvin!"},
	{"v2", "042", true, "Hello Marvin!"},
	{"v3", "042", true, "How do you do Marvin!"},
	{"v1", "666", false, "Go to hell Zuul!"},
	{"v2", "666", false, "Go to hell Zuul!"},
	{"v3", "666", false, "Go to hell Zuul!"},
}

func TestHello(t *testing.T) {

	for _, tt := range testHelloDatas {
		testHello(tt, t)
	}
}

func testHello(tData TestHelloData, t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"deserve\":\""+strconv.FormatBool(tData.deserve)+"\"}")
	}))
	defer ts.Close()
	os.Setenv("DESERVE_ADDR", ts.URL)
	logger := log.New(os.Stdout, "tst ", 0)
	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), logger, new(http.Client))
	handlers := greetings.NewHandlers(logger, deserveClient)
	r, _ := http.NewRequest("GET", "/itisnotchecked", nil)
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"version": tData.version,
		"userId":  tData.userId,
	})
	handlers.Hello().ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\""+tData.message+"\"}\n", string(w.Body.Bytes()))
}

func TestHello_UnsupportedVersion(t *testing.T) {

	logger := log.New(os.Stdout, "tst ", 0)
	handlers := greetings.NewHandlers(logger, nil)

	r, _ := http.NewRequest("GET", "/itisnotchecked", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"version": "v4",
		"userId":  "",
	})

	handlers.Hello().ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Unsupported Version v4\n", string(w.Body.Bytes()))

}
