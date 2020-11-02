package testutils

import (
	"net/http"
	"net/http/httptest"
)

// ServerMock create a specific mock server via parameter
// a api
// b response bytes
// c response code
func ServerMock(a string, b []byte, c int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(a, responseMocker(b, c))

	return httptest.NewServer(handler)
}

// ResponseMocker mock a specific response via parameters
// b response bytes
// c response code
func responseMocker(b []byte, c int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(c)
		w.Write(b)
	}
}
