package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

func TestEnvsNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	deploy, err := Envs("")

	assert.Nil(t, deploy)
	assert.Equal(t, `Get "http:///app//envs": http: no Host in request URL`, err.Error())
}

func TestEnvs(t *testing.T) {
	// init params
	appID := model.MockEnv.AppID
	// build mock server
	j, _ := json.Marshal(model.MockEnvsArray)
	s := testutils.ServerMock(fmt.Sprintf(EnvsAPI, appID), j, http.StatusOK)
	defer s.Close()

	// init stormkit
	viper.Set("app.server", s.URL[7:])
	stormkit.Config()

	envs, err := Envs(appID)

	assert.Equal(t, &model.MockEnvsArray, envs)
	assert.Nil(t, err)
}

func TestEnvs403(t *testing.T) {
	// init params
	appID := model.MockEnv.AppID
	// build mock server
	s := testutils.ServerMock(fmt.Sprintf(EnvsAPI, appID), nil, http.StatusForbidden)
	defer s.Close()

	// init stormkit
	viper.Set("app.server", s.URL[7:])
	stormkit.Config()

	// call API
	envs, err := Envs(appID)

	assert.Nil(t, envs)
	assert.Equal(t, `Error while doing request (response: 403 Forbidden)`, err.Error())
}
