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

func TestDeployByIDNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	deploy, err := DeployByID("", "")

	assert.Nil(t, deploy)
	assert.Equal(t, `Get "http:///app//deploy/": http: no Host in request URL`, err.Error())
}

func TestDeployByID(t *testing.T) {
	appID := model.MockSingleDeploy.Deploy.AppID
	id := model.MockSingleDeploy.Deploy.ID
	// build mock server
	j, _ := json.Marshal(model.MockSingleDeploy)
	s := testutils.ServerMock(fmt.Sprintf(DeployByIDapi, appID, id), j, http.StatusOK)
	defer s.Close()

	// set parameter and call API
	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()

	deploy, err := DeployByID(
		model.MockSingleDeploy.Deploy.AppID,
		model.MockSingleDeploy.Deploy.ID,
	)

	// test response
	assert.Nil(t, err)
	assert.Equal(t, &model.MockSingleDeploy, deploy)
}

func TestDeployByID403(t *testing.T) {
	appID := model.MockSingleDeploy.Deploy.AppID
	id := model.MockSingleDeploy.Deploy.ID

	// build mock server
	s := testutils.ServerMock(fmt.Sprintf(DeployByIDapi, appID, id), nil, http.StatusForbidden)
	defer s.Close()

	viper.Set("app.server", s.URL[7:])
	stormkit.Config()

	deploy, err := DeployByID(
		model.MockSingleDeploy.Deploy.AppID,
		model.MockSingleDeploy.Deploy.ID,
	)

	assert.Nil(t, deploy)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusForbidden))
}
