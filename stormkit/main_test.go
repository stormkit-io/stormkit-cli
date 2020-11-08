package stormkit

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"sync"
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
	Config("")

	// check Config runned correctly
	assert.Equal(t, expectedServer, globalConfig.Server)
	assert.Equal(t, expectedBearerToken, globalConfig.BearerToken)
	assert.Equal(t, time.Duration(expectedClientTimeout), globalConfig.ClientTimeout)
	assert.Equal(t, expectedUseHTTPS, globalConfig.UseHTTPS)
	assert.Equal(t, expectedEngineAppID, globalConfig.AppID)
}

// TestEngineAppID runs a sequence of manipulation on engineAppID
func TestEngineAppID(t *testing.T) {
	viperInit()

	Config("")

	assert.Equal(t, expectedEngineAppID, globalConfig.AppID)
	assert.Equal(t, expectedEngineAppID, GetEngineAppID())

	localEngineAppID := "test"

	SetEngineAppID(localEngineAppID)

	assert.Equal(t, localEngineAppID, globalConfig.AppID)
	assert.Equal(t, localEngineAppID, GetEngineAppID())
	assert.Equal(t, localEngineAppID, viper.Get(engineAppIDString))
}

func TestEngineAppIDConfigFile(t *testing.T) {
	viperInit()
	expectedAppID := "aaaa"
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  - id: ` + expectedAppID), nil
	}

	re := captureOutput(func() {
		Config(".")
	})

	assert.Equal(t, "", re)
}

func TestEngineAppIDConfigFileTwoApps(t *testing.T) {
	viperInit()
	expectedAppID := "aaaa"
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  - id: ` + expectedAppID + `
  - id: "ccccc"`), nil
	}

	re := captureOutput(func() {
		Config(".")
	})

	assert.Equal(t, "there are many apps in the config file (using the first)\n", re)
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
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

func TestReadStormkitConfigFailIO(t *testing.T) {
	expectedError := errors.New("error")
	ioutilReadFile = func(p string) ([]byte, error) {
		return nil, expectedError
	}

	configFile, err := readStormkitConfig("")

	assert.Nil(t, configFile)
	assert.Equal(t, expectedError, err)
}

func TestReadStormkitConfigUnmarshal(t *testing.T) {
	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  id: 10`), nil
	}

	configFile, err := readStormkitConfig("")

	assert.Nil(t, configFile)
	expectedErr := "yaml: unmarshal errors:\n  line 3: cannot unmarshal !!map into []struct { ID string \"yaml:\\\"id\\\"\" }"
	assert.Equal(t, expectedErr, err.Error())
}

func TestReadStormkitConfig(t *testing.T) {
	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  - id: 10`), nil
	}

	configFile, err := readStormkitConfig("")

	expectedConfigFile := ConfigFile{
		App: []struct {
			ID string `yaml:"id"`
		}{
			{
				ID: "10",
			},
		},
	}

	assert.Equal(t, &expectedConfigFile, configFile)
	assert.Nil(t, err)
}

func TestLoadStormkitConfigNoFile(t *testing.T) {
	expectedError := errors.New("error")
	osStat = func(p string) (os.FileInfo, error) {
		return nil, expectedError
	}

	err := loadStormkitConfig("")

	assert.Equal(t, expectedError, err)
}

// Fare test worng yml
func TestLoadStormkitConfigWrongYml(t *testing.T) {
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	expectedError := errors.New("error")
	ioutilReadFile = func(p string) ([]byte, error) {
		return nil, expectedError
	}

	err := loadStormkitConfig("")

	assert.Equal(t, expectedError, err)
}

func TestLoadStormkitConfigNoApp(t *testing.T) {
	expectedAppID := "aaaa"
	globalConfig.AppID = expectedAppID
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:`), nil
	}

	err := loadStormkitConfig("")

	assert.Nil(t, err)
	assert.Equal(t, expectedAppID, globalConfig.AppID)
}

func TestLoadStormkitConfigTwoApps(t *testing.T) {
	expectedAppID := "aaa"
	globalConfig.AppID = "bbb"
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  - id: ` + expectedAppID + `
  - id: 10`), nil
	}

	err := loadStormkitConfig("")

	assert.Equal(t, ErrMultipleAppsInConfigFile, err)
	assert.Equal(t, expectedAppID, globalConfig.AppID)
}

func TestLoadStormkitConfig(t *testing.T) {
	expectedAppID := "aaa"
	globalConfig.AppID = "bbb"
	fi := localFileInfo{}
	osStat = func(p string) (os.FileInfo, error) {
		return &fi, nil
	}
	fi.IsDirVar = false

	ioutilReadFile = func(p string) ([]byte, error) {
		return []byte(`
app:
  - id: ` + expectedAppID), nil
	}

	err := loadStormkitConfig("")

	assert.Nil(t, err)
	assert.Equal(t, expectedAppID, globalConfig.AppID)
}
