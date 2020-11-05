package stormkit

import (
	"errors"
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

// GlobalConfig is the configuration of the stomkit command
type GlobalConfig struct {
	Server        string        // address of the stormkit api server
	BearerToken   string        // access token to the api server
	ClientTimeout time.Duration // duration of the http client timeout
	UseHTTPS      bool          // enable https communication to the stormkit api server
	AppID         string        // id of the app to use
}

// osStat is the os.Stat abstraction function variable
var osStat = os.Stat

// ioutilReadFile is the ioutil.ReadFile abstraction function variable
var ioutilReadFile = ioutil.ReadFile

// globalConfig is the active general configuration
var globalConfig = GlobalConfig{}

// ErrManyAppInConfigFile error for many apps in config file
var ErrManyAppInConfigFile = errors.New("there are many apps in the config file (using the first)")

// Config configure the system for queries via viper (config file)
func Config() {
	globalConfig.Server = viper.GetString(serverString)
	globalConfig.BearerToken = viper.GetString(bearerTokenString)
	globalConfig.ClientTimeout = time.Duration(viper.GetInt64(clientTimeoutString))
	globalConfig.UseHTTPS = viper.GetBool(useHTTPSString)
	globalConfig.AppID = viper.GetString(engineAppIDString)
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
	return globalConfig.AppID
}

// SetEngineAppID set the value in engineAppID and in the stormkit file
func SetEngineAppID(a string) error {
	globalConfig.AppID = a
	viper.Set(engineAppIDString, a)
	return viper.WriteConfig()
}

// GetGlobalConfig return the global configuration
func GetGlobalConfig() *GlobalConfig {
	return &globalConfig
}
