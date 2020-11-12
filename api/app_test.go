package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

// expectedApps is mock apps data for tests
var ExpectedApps = Apps{
	Apps: []App{
		{
			Repo: "repo1",
			ID:   "1234",
		},
		{
			Repo: "repo2",
			ID:   "12345",
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
	j, _ := json.Marshal(ExpectedApps)
	s := testutils.ServerMock("/apps", j, http.StatusOK)
	defer s.Close()

	// set parameters and call API
	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	// test responses
	assert.Nil(t, err)
	assert.Equal(t, &ExpectedApps, apps)
}

func TestGetApps403(t *testing.T) {
	s := testutils.ServerMock("/apps", nil, http.StatusForbidden)
	defer s.Close()

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()
	apps, err := GetApps()

	assert.Nil(t, apps)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusForbidden))
}

func TestDumpApp(t *testing.T) {
	s := DumpApp(ExpectedApps.Apps[0])
	createdAt := time.Unix(int64(ExpectedApps.Apps[0].CreatedAt), 0)
	deployedAt := time.Unix(int64(ExpectedApps.Apps[0].DeployedAt), 0)
	expected := fmt.Sprintf("Repo: repo1\n  ID: 1234\n  Status: false\n  AutoDeploy: \n  DefaultEnv: \n  Endpoint: \n  DisplayName: \n  CreatedAt: %s\n  DeployedAt: %s\n\n", createdAt, deployedAt)

	assert.Equal(t, expected, s)
}
