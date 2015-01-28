package semaphore

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func testApiCall(code int, body string) (*httptest.Server, *Api) {
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
	api := &Api{"123", server.URL, httpClient}

	return server, api
}
