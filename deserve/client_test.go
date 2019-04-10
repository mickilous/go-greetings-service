package deserve_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"greetings-service/configuration"
	"greetings-service/deserve"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIsGreetable_False(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"deserve\":\"false\"}")
	}))
	defer ts.Close()

	os.Setenv("DESERVE_ADDR", ts.URL)

	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), log.New(os.Stdout, "tst ", 0), new(http.Client))
	greetable := deserveClient.IsGreetable("042")
	assert.False(t, greetable, "Greetable must be false")

}

func TestIsGreetable_HTTP_protocol_error(t *testing.T) {

	os.Clearenv()

	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), log.New(os.Stdout, "tst ", 0), new(http.Client))
	greetable := deserveClient.IsGreetable("042")
	assert.True(t, greetable, "Greetable must be true (fallback)")

}

func ExampleIsGreetable_HTTP_protocol_error() {

	os.Clearenv()

	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), log.New(os.Stdout, "tst ", 0), new(http.Client))
	deserveClient.IsGreetable("042")
	// Output:
	// tst The HTTP request failed with error Get http://localhost:8090/deserve/042: dial tcp [::1]:8090: connect: connection refused

}

//func TestIsGreetable_HTTP_body_read_error(t *testing.T) {
//
//
//}

func TestIsGreetable_HttpStatus_Error(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "This is not good")
	}))
	defer ts.Close()

	os.Setenv("DESERVE_ADDR", ts.URL)

	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), log.New(os.Stdout, "tst ", 0), new(http.Client))
	greetable := deserveClient.IsGreetable("042")
	assert.True(t, greetable, "Greetable must be true (fallback)")

}

func ExampleIsGreetable_HttpStatus_Error() {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "This is not good")
	}))
	defer ts.Close()

	os.Setenv("DESERVE_ADDR", ts.URL)

	deserveClient := deserve.NewClient(configuration.NewEnvironmentProvider(), log.New(os.Stdout, "tst ", 0), new(http.Client))
	deserveClient.IsGreetable("042")
	// Output:
	// tst The HTTP request failed with HTTP Code : 400 - This is not good

}
