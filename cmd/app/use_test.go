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
	cmd.Flags().String("app-id", "", "")

	return s, &cmd
}

func TestRunAppUseNoFlag(t *testing.T) {
	s, _ := runAppUseInit()
	defer s.Close()

	cmd := cobra.Command{}
	args := []string{}
	err := runAppUse(&cmd, args)

	assert.Equal(t, "flag accessed but not defined: app-id", err.Error())
}

func TestRunAppUseAppId(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	cmd.Flags().Set("app-id", expectedApps.Apps[0].ID)
	args := []string{}
	err := runAppUse(cmd, args)

	assert.Nil(t, err)
	assert.Equal(t, expectedApps.Apps[0].ID, viper.GetString("app.engine.app_id"))
}

func TestRunAppUseAppIDNotFound(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	cmd.Flags().Set("app-id", "1")
	args := []string{}
	err := runAppUse(cmd, args)

	assert.Equal(t, "no app found", err.Error())
}

func TestRunAppUseNotEnoughtArgs(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{}
	err := runAppUse(cmd, args)

	assert.Equal(t, "not enought arguments", err.Error())
}

func TestRunAppUseRepo(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{expectedApps.Apps[0].Repo}
	err := runAppUse(cmd, args)

	assert.Nil(t, err)
	assert.Equal(t, expectedApps.Apps[0].ID, viper.GetString("app.engine.app_id"))
}

func TestRunAppUseRepoNotFound(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{"a"}
	err := runAppUse(cmd, args)

	assert.Equal(t, "app not found", err.Error())
}

func TestRunAppUseIndexTooLarge(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{"2"}
	err := runAppUse(cmd, args)

	assert.Equal(t, "index too large", err.Error())
}

func TestRunAppUse(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()

	args := []string{"0"}
	err := runAppUse(cmd, args)

	assert.Nil(t, err)
	assert.Equal(t, expectedApps.Apps[0].ID, viper.Get("app.engine.app_id"))
}
