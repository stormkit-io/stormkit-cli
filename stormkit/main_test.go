package stormkit

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	expectedServer := "aaaa"
	viper.Set(serverString, expectedServer)
	expectedBearerToken := "bbbbb"
	viper.Set(bearerTokenString, expectedBearerToken)

	Config()

	assert.Equal(t, expectedServer, server)
	assert.Equal(t, expectedBearerToken, bearerToken)
}
