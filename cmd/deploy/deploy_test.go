package deploy

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

// mockPromptSelect is mock implementation of promptSelect interface
// for testing of function with promptui.Select
type mockPromptSelect struct {
	I int
	S string
	E error
}

// Run method of mockPromptSelect for implement promptSelect
func (p *mockPromptSelect) Run() (int, string, error) {
	return p.I, p.S, p.E
}

// TestRunEnvPromptError execute tessts with promptSelect.Run()
// giving an error
func TestRunEnvPromptError(t *testing.T) {
	// save envPrompt original function
	ep := envPrompt

	// prepare mockPromptSelect and expected data
	m := mockPromptSelect{}
	m.I = 1
	m.S = ""
	m.E = fmt.Errorf("error")
	expectedI := -1
	expectedError := m.E
	envPrompt = func(*model.EnvsArray) promptSelect {
		return &m
	}

	// execute runEnvPrompt
	i, err := runEnvPrompt(nil)

	// check tests results
	assert.Equal(t, expectedI, i)
	assert.Equal(t, expectedError, err)

	// restore original envPrompt function
	envPrompt = ep
}

// TestRunEnvPrompt execute tests with promptSelect.Run()
// retriving no errors
func TestRunEnvPrompt(t *testing.T) {
	// save envPrompt original function
	ep := envPrompt

	// prepare mockPromptSelect and expected data
	m := mockPromptSelect{}
	m.I = 1
	m.S = ""
	m.E = nil
	envPrompt = func(*model.EnvsArray) promptSelect {
		return &m
	}

	// execute runEnvPrompt
	i, err := runEnvPrompt(nil)

	// check test results
	assert.Equal(t, m.I, i)
	assert.Equal(t, m.E, err)

	// restore original envPrompt function
	envPrompt = ep
}

func TestDeployInteractiveNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	d, err := deployInteractive()

	assert.Nil(t, d)
	assert.Equal(t, `Get "http:///app/12346/envs": http: no Host in request URL`, err.Error())
}
