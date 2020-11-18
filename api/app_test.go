package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

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
	j, _ := json.Marshal(model.MockApps)
	s := testutils.ServerMock(API.Apps, j, http.StatusOK)
	defer s.Close()

	// set parameters and call API
	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	// test responses
	assert.Nil(t, err)
	assert.Equal(t, &model.MockApps, apps)
}

func TestGetApps403(t *testing.T) {
	s := testutils.ServerMock(API.Apps, nil, http.StatusForbidden)
	defer s.Close()

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	assert.Nil(t, apps)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusForbidden))
}
