package stormkit

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const expectedServer = "aaaa"
const expectedBearerToken = "bbbb"
const expectedClientTimeout = 5000
const expectedUseHTTPS = true

// viperInit initialize viper for testing
func viperInit() {
	viper.Set(serverString, expectedServer)
	viper.Set(bearerTokenString, expectedBearerToken)
	viper.Set(clientTimeoutString, expectedClientTimeout)
	viper.Set(useHTTPSString, expectedUseHTTPS)
}

func TestConfig(t *testing.T) {
	viperInit()

	// run Config
	Config()

	// check Config runned correctly
	assert.Equal(t, expectedServer, server)
	assert.Equal(t, expectedBearerToken, bearerToken)
	assert.Equal(t, time.Duration(expectedClientTimeout), clientTimeout)
	assert.Equal(t, expectedUseHTTPS, useHTTPS)
}
