package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

var expectedApps = api.Apps{
	Apps: []api.App{
		{ID: "123", Repo: "repo0"},
		{ID: "124", Repo: "repo1"},
	},
}

func runAppUseInit() (*httptest.Server, *cobra.Command) {
	j, _ := json.Marshal(expectedApps)
	s := testutils.ServerMock("/apps", j, http.StatusOK)

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()

	cmd := cobra.Command{}

	return s, &cmd
}

func TestRunAppUseNotServer(t *testing.T) {
	viper.Set("app.server", "aaaa")
	stormkit.Config()

	cmd := cobra.Command{}
	args := []string{}
	err := runAppUse(&cmd, args)

	assert.Equal(t, `Get http://aaaa/apps: dial tcp: lookup aaaa: no such host`, err.Error())
}

func TestRunAppUseNotFound(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{"1"}
	err := runAppUse(cmd, args)

	assert.Equal(t, "app not found", err.Error())
}

func TestRunAppUseNotEnoughtArgs(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{}
	err := runAppUse(cmd, args)

	assert.Equal(t, "not enought arguments", err.Error())
}

func TestRunAppUse(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{expectedApps.Apps[0].ID}
	err := runAppUse(cmd, args)

	assert.Nil(t, err)
	assert.Equal(t, expectedApps.Apps[0].ID, viper.Get("app.engine.app_id"))
}
