package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

func runAppUseInit() (*httptest.Server, *cobra.Command) {
	j, _ := json.Marshal(model.MockApps)
	s := testutils.ServerMock("/apps", j, http.StatusOK)

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()

	cmd := cobra.Command{}

	return s, &cmd
}

func TestRunAppUseNotServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := cobra.Command{}
	args := []string{}
	err := runAppUse(&cmd, args)

	assert.Equal(t, `Get "http:///apps": http: no Host in request URL`, err.Error())
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

	args := []string{model.MockApp.ID}
	err := runAppUse(cmd, args)

	assert.Nil(t, err)
	assert.Equal(t, model.MockApp.ID, viper.Get("app.engine.app_id"))
}
