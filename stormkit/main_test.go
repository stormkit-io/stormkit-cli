package stormkit

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// setting data in viper
	expectedServer := "aaaa"
	viper.Set(serverString, expectedServer)

	expectedBearerToken := "bbbbb"
	viper.Set(bearerTokenString, expectedBearerToken)

	expectedClientTimeout := 5000
	viper.Set(clientTimeoutString, expectedClientTimeout)

	expectedUseHTTPS := true
	viper.Set(useHTTPSString, expectedUseHTTPS)

	// run Config
	Config()

	// check Config runned correctly
	assert.Equal(t, expectedServer, server)
	assert.Equal(t, expectedBearerToken, bearerToken)
	assert.Equal(t, time.Duration(expectedClientTimeout), clientTimeout)
	assert.Equal(t, expectedUseHTTPS, useHTTPS)
}
