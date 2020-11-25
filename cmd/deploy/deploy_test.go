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

// mockPrompt is mock implementation of prompt interface for
// testing of function with promptui.Prompt
type mockPrompt struct {
	S string
	E error
}

// Run method of mockPromptSelect for implement promptSelect
func (p *mockPrompt) Run() (string, error) {
	return p.S, p.E
}

func TestRunBranchPrompt(t *testing.T) {
	// save branchPrompt original function
	bp := branchPrompt

	// prepare mockPrompt and expected data
	stormkit.RepoPath = ""
	m := mockPrompt{}
	m.S = "helo"
	m.E = nil
	branchPrompt = func() prompt {
		return &m
	}

	// execute runBranchPrompt
	b, err := runBranchPrompt(m.S)

	// check test result
	assert.Equal(t, m.S, b)
	assert.Equal(t, m.E, err)

	// restore original branchPrompt function
	branchPrompt = bp
}

func TestRunBranchPromptError(t *testing.T) {
	// save branchPrompt original function
	bp := branchPrompt

	// prepare mockData
	stormkit.RepoPath = ""
	m := mockPrompt{}
	m.S = ""
	m.E = fmt.Errorf("error")
	branchPrompt = func() prompt {
		return &m
	}

	// execute runBranchPrompt
	b, err := runBranchPrompt(m.S)

	// check test result
	assert.Equal(t, m.S, b)
	assert.Equal(t, m.E, err)

	// restore ogiginal branch Prompt
	branchPrompt = bp
}

// mockSelectWithAdd is mock implementation of prompt interface for
// testing of function with promptui.Prompt
type mockSelectWithAdd struct {
	I int
	S string
	E error
}

// Run method of mockSelectWithAdd for implement promptSelect
func (p *mockSelectWithAdd) Run() (int, string, error) {
	return p.I, p.S, p.E
}

func TestRunBranchPromptSelectError(t *testing.T) {
	// save branchPrompt original function
	bp := branchSelectWithAdd

	// prepare mockData
	stormkit.RepoPath = "./../.."
	m := mockSelectWithAdd{}
	m.I = -1
	m.S = ""
	m.E = fmt.Errorf("error")
	branchSelectWithAdd = func(b []string) selectWithAdd {
		return &m
	}

	// execute runBranchPrompt
	b, err := runBranchPrompt(m.S)

	// check test result
	assert.Equal(t, "", b)
	assert.Equal(t, m.E, err)

	// restore ogiginal branch Prompt
	branchSelectWithAdd = bp
}

func TestRunBranchPromptSelectDefault(t *testing.T) {
	// save branchPrompt original function
	bp := branchSelectWithAdd

	// prepare mockData
	stormkit.RepoPath = "./../.."
	m := mockSelectWithAdd{}
	m.I = 0
	m.S = ""
	m.E = nil
	branchSelectWithAdd = func(b []string) selectWithAdd {
		return &m
	}
	expectedBranch := "branch"

	// execute runBranchPrompt
	b, err := runBranchPrompt(expectedBranch)

	// check test result
	assert.Equal(t, expectedBranch, b)
	assert.Nil(t, err)

	// restore ogiginal branch Prompt
	branchSelectWithAdd = bp
}

func TestRunBranchPromptSelect(t *testing.T) {
	// save branchPrompt original function
	bp := branchSelectWithAdd

	// prepare mockData
	stormkit.RepoPath = "./../.."
	m := mockSelectWithAdd{}
	m.I = 1
	m.S = "helo"
	m.E = nil
	branchSelectWithAdd = func(b []string) selectWithAdd {
		return &m
	}

	// execute runBranchPrompt
	b, err := runBranchPrompt("hy")

	// check test result
	assert.Equal(t, m.S, b)
	assert.Nil(t, err)

	// restore ogiginal branch Prompt
	branchSelectWithAdd = bp
}

func TestDeployInteractiveNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	d, err := deployInteractive()

	assert.Nil(t, d)
	assert.Equal(t, `Get "http:///app/12346/envs": http: no Host in request URL`, err.Error())
}
