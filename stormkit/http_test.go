package stormkit

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	viperInit()

	Config("")

	c := GetClient()

	assert.Equal(t, time.Duration(expectedClientTimeout), c.Timeout)
}

func TestRequest(t *testing.T) {
	viperInit()
	Config("")

	r, err := request("GET", "/api", nil)
	auth := r.Header.Get(authorizationHeaderString)
	assert.Contains(t, auth, expectedBearerToken)
	assert.Nil(t, err)
	assert.Equal(t, "GET", r.Method)
}

func TestRequestGet(t *testing.T) {
	viperInit()
	Config("")

	r, err := Get("a")

	assert.Equal(t, http.MethodGet, r.Method)
	assert.Nil(t, err)
}

func TestRequestPost(t *testing.T) {
	viperInit()
	Config("")

	r, err := Post("a", nil)

	assert.Equal(t, http.MethodPost, r.Method)
	assert.Nil(t, err)
}

func TestRequestPut(t *testing.T) {
	viperInit()
	Config("")

	r, err := Put("a", nil)

	assert.Equal(t, http.MethodPut, r.Method)
	assert.Nil(t, err)
}

func TestRequestDelete(t *testing.T) {
	viperInit()
	Config("")

	r, err := Delete("a", nil)

	assert.Equal(t, http.MethodDelete, r.Method)
	assert.Nil(t, err)
}
