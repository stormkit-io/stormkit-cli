package stormkit

import (
	"time"

	"github.com/spf13/viper"
)

const (
	serverString = "app.server" // is the key of server configuration
	bearerTokenString = "app.bearer_token" // is the key of bearer token configuration
	clientTimeoutString = "app.client_timeout" // is the timeout of the http client
)

// server is the address to the server
var server string
// bearerToken is used to access http
var bearerToken string
// clientTimeout is the timeout of the http client
var clientTimeout time.Duration

// Config configure the system for queries via viper (config file)
func Config() {
	server = viper.GetString(serverString)
	bearerToken = viper.GetString(bearerTokenString)
	clientTimeout = time.Duration(viper.GetInt64(clientTimeoutString))
}

