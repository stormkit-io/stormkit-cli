package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
)

// expectedApps is mock apps data for tests
var expectedApps = Apps{
	Apps: []App{
		{
			Repo: "repo1",
		},
		{
			Repo: "repo2",
		},
	},
}

// TestGetApps no app.server parameter
func TestGetAppsNoServer(t *testing.T) {
	stormkit.Config()
	apps, err := GetApps()

	assert.Nil(t, apps)
	assert.NotNil(t, err)
}

// TestGetApps with http code OK
func TestGetApps(t *testing.T) {
	// build mock server
	j, _ := json.Marshal(expectedApps)
	s := serverMock("/apps", j, http.StatusOK)
	defer s.Close()

	// set parameters and call API
	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	// test responses
	assert.Nil(t, err)
	assert.Equal(t, &expectedApps, apps)
}

func TestGetApps403(t *testing.T) {
	s := serverMock("/apps", nil, http.StatusForbidden)
	defer s.Close()

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	assert.Nil(t, apps)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusForbidden))
}

// serverMock create a specific mock server via parameter
// a api
// b response bytes
// c response code
func serverMock(a string, b []byte, c int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(a, responseMocker(b, c))

	return httptest.NewServer(handler)
}

// responseMocker mock a specific response via parameters
// b response bytes
// c response code
func responseMocker(b []byte, c int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(c)
		w.Write(b)
	}
}
