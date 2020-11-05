package stormkit

import (
	"os"
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

func TestGetStormkitConfigFilePathNotExits(t *testing.T) {
	osStat = func(p string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}

	p, err := getStormkitConfigFilePath("")

	assert.Empty(t, p)
	assert.Equal(t, os.ErrNotExist, err)
}

type localFileInfo struct {
	NameVar    string
	SizeVar    int64
	ModeVar    os.FileMode
	ModTimeVar time.Time
	IsDirVar   bool
	SysVar     interface{}
}

func (fi *localFileInfo) Name() string {
	return fi.NameVar
}

func (fi *localFileInfo) Size() int64 {
	return fi.SizeVar
}

func (fi *localFileInfo) Mode() os.FileMode {
	return fi.ModeVar
}

func (fi *localFileInfo) ModTime() time.Time {
	return fi.ModTimeVar
}

func (fi *localFileInfo) IsDir() bool {
	return fi.IsDirVar
}

func (fi *localFileInfo) Sys() interface{} {
	return fi.SysVar
}

func TestGetStormkitConfigFileDir(t *testing.T) {
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = true

	p, err := getStormkitConfigFilePath("")

	assert.Empty(t, p)
	assert.Equal(t, "/stormkit.config.yml is a directory not a file", err.Error())
}

func TestGetStormkitConfigFile(t *testing.T) {
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	path := "path"
	expectedPath := path + "/stormkit.config.yml"
	p, err := getStormkitConfigFilePath(path)

	assert.Equal(t, expectedPath, p)
	assert.Nil(t, err)
}
