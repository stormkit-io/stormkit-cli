package deploy

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

func prepareDeployCmd() *cobra.Command {
	c := cobra.Command{}
	c.Flags().Bool(interactiveFlag, false, "")

	return &c
}

func runDeployInit() (*httptest.Server, *cobra.Command) {
	j, _ := json.Marshal(model.MockDeploy)
	s := testutils.ServerMock(api.API.Deploy, j, http.StatusOK)

	viper.Set("app.server", s.URL[7:])
	viper.Set("app.engine.app_id", model.MockDeploy.AppID)
	stormkit.Config()

	return s, prepareDeployCmd()
}

func TestRunDeployNoFlag(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := cobra.Command{}
	args := []string{}
	err := runDeploy(&cmd, args)

	assert.Equal(t, "flag accessed but not defined: interactive", err.Error())
}

func TestRunDeployNoArgs(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := prepareDeployCmd()
	args := []string{}
	err := runDeploy(cmd, args)

	assert.Equal(t, "not enought arguments", err.Error())
}

func TestRunDeployNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := prepareDeployCmd()
	args := []string{"prod", "master"}
	err := runDeploy(cmd, args)

	assert.Equal(t, `Post "http:///app/deploy": http: no Host in request URL`, err.Error())
}

func TestRunDeploy(t *testing.T) {
	s, cmd := runDeployInit()
	defer s.Close()

	args := []string{"env", "branch"}
	f := func() {
		err := runDeploy(cmd, args)
		assert.Nil(t, err)
	}

	output := testutils.CaptureOutput(f)
	expectedOutput := "Deploy ID: " + model.MockDeploy.ID + "\n"

	assert.Equal(t, expectedOutput, output)

}

func TestDeployInteractiveNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	d, err := deployInteractive()

	assert.Nil(t, d)
	assert.Equal(t, `Get "http:///app/12346/envs": http: no Host in request URL`, err.Error())
}
