package stormkit

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	serverString        = "app.server"         // is the key of server configuration
	bearerTokenString   = "app.bearer_token"   // is the key of bearer token configuration
	clientTimeoutString = "app.client_timeout" // is the timeout of the http client
	useHTTPSString      = "app.https"          // is the flag for use https
	engineAppIDString   = "app.engine.app_id"  // is the place for store the app_id
)

// ConfigFile is the rappresentation of the stormkit config file `stormkit.config.yml`
type ConfigFile struct {
	App []struct {
		ID string `yaml:"id"`
	}
}

// server is the address to the server
var server string

// bearerToken is used to access http
var bearerToken string

// clientTimeout is the timeout of the http client
var clientTimeout time.Duration

// useHTTPS is the flag for use https in http requests
var useHTTPS bool

// engineAppID is the place for store the active app_id
var engineAppID string

// osStat is the os.Stat abstraction function variable
var osStat = os.Stat

// ioutilReadFile is the ioutil.ReadFile abstraction function variable
var ioutilReadFile = ioutil.ReadFile

// Config configure the system for queries via viper (config file)
func Config() {
	server = viper.GetString(serverString)
	bearerToken = viper.GetString(bearerTokenString)
	clientTimeout = time.Duration(viper.GetInt64(clientTimeoutString))
	useHTTPS = viper.GetBool(useHTTPSString)
	engineAppID = viper.GetString(engineAppIDString)
}

// getStormkitConfigFilePath check if in the folder is a stormkit config file
// checks before the stormkit.config.yml then stormkit.config.yaml
func getStormkitConfigFilePath(repoPath string) (string, error) {
	path := repoPath + "/stormkit.config.yml"

	info, err := osStat(path)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("%s is a directory not a file", path)
	}

	return path, nil
}

// readStormkitConfig read the stormkit config file and places it in a
// ConfigFile struct
func readStormkitConfig(path string) (*ConfigFile, error) {
	// read config file via ioutil.ReadFile abstraction
	ymlFile, err := ioutilReadFile(path)
	if err != nil {
		return nil, err
	}

	// map config file to struct
	var c *ConfigFile
	err = yaml.Unmarshal(ymlFile, &c)
	if err != nil {
		return nil, err
	}

	return c, err
}

// GetEngineAppID return the value of engineAppID
func GetEngineAppID() string {
	return engineAppID
}

// SetEngineAppID set the value in engineAppID and in the stormkit file
func SetEngineAppID(a string) error {
	engineAppID = a
	viper.Set(engineAppIDString, a)
	return viper.WriteConfig()
}
