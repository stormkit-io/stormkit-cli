package log

import (
	"encoding/json"
	"fmt"
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

func runLogInit() (*httptest.Server, *cobra.Command) {
	j, _ := json.Marshal(model.MockSingleDeploy)
	apiurl := fmt.Sprintf(api.API.DeployByID, model.MockDeploy.AppID, model.MockDeploy.ID)
	s := testutils.ServerMock(apiurl, j, http.StatusOK)

	viper.Set("app.server", s.URL[7:])
	viper.Set("app.engine.app_id", model.MockDeploy.AppID)
	stormkit.Config()

	cmd := cobra.Command{}

	return s, &cmd
}

func TestRunLogNoArgs(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := cobra.Command{}
	args := []string{}
	err := runLog(&cmd, args)

	assert.Equal(t, "not enought arguments", err.Error())
}

func TestRunLogNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	cmd := cobra.Command{}
	args := []string{"12345"}
	err := runLog(&cmd, args)

	assert.Equal(t, `Get "http:///app//deploy/12345": http: no Host in request URL`, err.Error())
}

func TestRunLogNotFound(t *testing.T) {
	s, cmd := runLogInit()
	defer s.Close()

	args := []string{"1"}
	err := runLog(cmd, args)

	assert.Equal(t, "Error while doing request (response: 404 Not Found)", err.Error())
}

func TestRunLog(t *testing.T) {
	s, cmd := runLogInit()
	defer s.Close()

	args := []string{model.MockDeploy.ID}

	f := func() {
		err := runLog(cmd, args)
		assert.Nil(t, err)
	}

	output := testutils.CaptureOutput(f)
	expectedOutput := ""
	for _, l := range model.MockDeploy.Logs {
		expectedOutput += l.Dump() + "\n"
	}

	assert.Equal(t, expectedOutput, output)
}
