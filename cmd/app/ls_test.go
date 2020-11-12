package app

import (
	"encoding/json"
	"fmt"
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

func runAppLsInit() (*httptest.Server, *cobra.Command) {
	j, _ := json.Marshal(expectedApps)
	s := testutils.ServerMock("/apps", j, http.StatusOK)

	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()

	cmd := cobra.Command{}
	cmd.Flags().Bool("details", false, "")

	return s, &cmd
}

func TestRunAppLsNotServer(t *testing.T) {
	stormkit.Config()

	viper.Set("app.server", "")
	cmd := cobra.Command{}
	args := []string{}

	err := runAppLs(&cmd, args)

	assert.Equal(t, `Get "http:///apps": http: no Host in request URL`, err.Error())
}

func TestRunAppLsNoFlag(t *testing.T) {
	s, cmd := runAppUseInit()
	defer s.Close()
	args := []string{}

	err := runAppLs(cmd, args)

	assert.Equal(t, `flag accessed but not defined: details`, err.Error())
}

func TestRunAppLs(t *testing.T) {
	s, cmd := runAppLsInit()
	defer s.Close()
	args := []string{}

	f := func() {
		err := runAppLs(cmd, args)
		assert.Nil(t, err)
	}

	out := testutils.CaptureOutput(f)

	expectedOutput := fmt.Sprintf(
		"ID   Repository\n%s  %s\n%s  %s\n",
		expectedApps.Apps[0].ID,
		expectedApps.Apps[0].Repo,
		expectedApps.Apps[1].ID,
		expectedApps.Apps[1].Repo,
	)
	assert.Equal(t, expectedOutput, out)
}

func TestRunAppLsDetails(t *testing.T) {
	s, cmd := runAppLsInit()
	defer s.Close()
	args := []string{}
	cmd.Flags().Set("details", "true")

	f := func() {
		err := runAppLs(cmd, args)
		assert.Nil(t, err)
	}

	out := testutils.CaptureOutput(f)

	a := expectedApps.Apps[0]
	a0 := api.DumpApp(a)
	a = expectedApps.Apps[1]
	a1 := api.DumpApp(a)

	assert.Equal(t, a0+a1, out)
}
