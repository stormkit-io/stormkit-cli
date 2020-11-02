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
const expectedEngineAppID = "1010101010"

// viperInit initialize viper for testing
func viperInit() {
	viper.Set(serverString, expectedServer)
	viper.Set(bearerTokenString, expectedBearerToken)
	viper.Set(clientTimeoutString, expectedClientTimeout)
	viper.Set(useHTTPSString, expectedUseHTTPS)
	viper.Set(engineAppIDString, expectedEngineAppID)
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
	assert.Equal(t, expectedEngineAppID, engineAppID)
}

// TestEngineAppID runs a sequence of manipulation on engineAppID
func TestEngineAppID(t *testing.T) {
	viperInit()

	Config()

	assert.Equal(t, expectedEngineAppID, engineAppID)
	assert.Equal(t, expectedEngineAppID, GetEngineAppID())
	
	localEngineAppID := "test"

	SetEngineAppID(localEngineAppID)

	assert.Equal(t, localEngineAppID, engineAppID)
	assert.Equal(t, localEngineAppID, GetEngineAppID())
	assert.Equal(t, localEngineAppID, viper.Get(engineAppIDString))
}
