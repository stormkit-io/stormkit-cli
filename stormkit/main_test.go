package stormkit

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	expectedServer := "aaaa"
	viper.Set(serverString, expectedServer)
	expectedBearerToken := "bbbbb"
	viper.Set(bearerTokenString, expectedBearerToken)
	expectedClientTimeout := 5000
	viper.Set(clientTimeoutString, expectedClientTimeout)

	Config()

	assert.Equal(t, expectedServer, server)
	assert.Equal(t, expectedBearerToken, bearerToken)
	assert.Equal(t, time.Duration(expectedClientTimeout), clientTimeout)
}
