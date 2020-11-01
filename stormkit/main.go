package stormkit

import (
	"github.com/spf13/viper"
)

const (
	serverString = "app.server" // is the key of server in the configuration
	bearerTokenString = "app.bearer_token" // is the key of bearer token in the configuration
)

// server is the address to the server
var server string
// bearerToken is used to access http
var bearerToken string

// Config configure the system for queries via viper (config file)
func Config() {
	server = viper.GetString(serverString)
	bearerToken = viper.GetString(bearerTokenString)
}
