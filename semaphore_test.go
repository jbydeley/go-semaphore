package semaphore

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	api := &Semaphore{"123", server.URL, httpClient}

	return server, api
}

// To create an API you'll need to give your
// authorization token.
func ExampleExamples() {
	NewSemaphore("auth_token")
}
