package stormkit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	viperInit()

	Config()

	c := GetClient()

	assert.Equal(t, time.Duration(expectedClientTimeout), c.Timeout)
}

func TestRequest(t *testing.T) {
	viperInit()
	Config()

	r, err := request("GET", "/api", nil)
	auth := r.Header.Get(authorizationHeaderString)
	assert.Contains(t, auth, expectedBearerToken)
	assert.Nil(t, err)
	assert.Equal(t, "GET", r.Method)
}
