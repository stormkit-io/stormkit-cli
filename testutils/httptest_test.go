package testutils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// customResponseWriter is a quick implementation of ResponseWriter
type customResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

// Header returns the header of the customResponseWriter
func (w *customResponseWriter) Header() http.Header {
	return w.header
}

// Write write data in byte variable
func (w *customResponseWriter) Write(b []byte) (int, error) {
	w.body = b
	return 0, nil
}

// WriteHeader writes the statuCode
func (w *customResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// TestResponseMocker check the responseMocker method
func TestResponseMocker(t *testing.T) {
	expectedBytes := []byte{0, 0, 0}
	expectedCode := 200
	f := responseMocker(expectedBytes, expectedCode)

	w := customResponseWriter{}
	r := http.Request{}

	f(&w, &r)

	assert.Equal(t, expectedBytes, w.body)
	assert.Equal(t, expectedCode, w.statusCode)
}

func TestServerMock(t *testing.T) {
	expectedBytes := []byte{10, 20, 30, 40}
	expectedCode := 200
	api := "/api"

	s := ServerMock(api, expectedBytes, expectedCode)
	defer s.Close()

	fmt.Println(s.URL + api)
	request, err := http.NewRequest(http.MethodGet, s.URL+api, bytes.NewReader(expectedBytes))
	assert.Nil(t, err)
	c := http.Client{}

	response, err := c.Do(request)
	assert.Nil(t, err)

	assert.Equal(t, expectedCode, response.StatusCode)
	actualBytes, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBytes, actualBytes)
}
