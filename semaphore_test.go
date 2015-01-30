package semaphore

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func testAPICall(code int, body string) (*httptest.Server, *Semaphore) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	api := &Semaphore{"123", server.URL, httpClient, logger}

	return server, api
}

func Test_NewSemaphore(t *testing.T) {
	api := NewSemaphore("123")

	if api.authToken != "123" {
		t.Errorf("Auth token should have been 123 but was %v", api.authToken)
	}
}

func Test_NewSemaphoreWithLogger(t *testing.T) {
	logger := log.New(ioutil.Discard, "TEST", log.LstdFlags)
	api := NewSemaphoreWithLogger("123", logger)

	if api.authToken != "123" {
		t.Errorf("Auth token should have been 123 but was %v", api.authToken)
	}

	if !reflect.DeepEqual(logger, api.log) {
		t.Error("Logger didn't set correctly")
	}
}

// To create an API you'll need to give your
// authorization token.
func ExampleExamples() {
	NewSemaphore("auth_token")
}
