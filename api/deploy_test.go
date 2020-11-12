package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/testutils"
	"github.com/stretchr/testify/assert"
)

var expectedSingleDeploy = SingleDeploy{
	Deploy: Deploy{
		ID:    "12345",
		AppID: "12346",
		Logs: []Log{
			{
				Title:   "title log 0",
				Message: "message log 0 that is true",
				Status:  true,
			},
			{
				Title:   "title log 1",
				Message: "message log 1 that is false",
				Status:  false,
			},
		},
	},
}

func TestDeployByIDNoServer(t *testing.T) {
	viper.Set("app.server", "")
	stormkit.Config()

	deploy, err := DeployByID("", "")

	assert.Nil(t, deploy)
	assert.Equal(t, `Get "http:///app//deploy/": http: no Host in request URL`, err.Error())
}

func TestDeployByID(t *testing.T) {
	appID := expectedSingleDeploy.Deploy.AppID
	id := expectedSingleDeploy.Deploy.ID
	// build mock server
	j, _ := json.Marshal(expectedSingleDeploy)
	s := testutils.ServerMock(fmt.Sprintf(deployByIDapi, appID, id), j, http.StatusOK)
	defer s.Close()

	// set parameter and call API
	viper.Set("app.server", s.URL[7:len(s.URL)])
	stormkit.Config()

	deploy, err := DeployByID(
		expectedSingleDeploy.Deploy.AppID,
		expectedSingleDeploy.Deploy.ID,
	)

	// test response
	assert.Nil(t, err)
	assert.Equal(t, &expectedSingleDeploy, deploy)
}

func TestDeployByID403(t *testing.T) {
	appID := expectedSingleDeploy.Deploy.AppID
	id := expectedSingleDeploy.Deploy.ID

	// build mock server
	s := testutils.ServerMock(fmt.Sprintf(deployByIDapi, appID, id), nil, http.StatusForbidden)
	defer s.Close()

	viper.Set("app.server", s.URL[7:])
	stormkit.Config()

	deploy, err := DeployByID(
		expectedSingleDeploy.Deploy.AppID,
		expectedSingleDeploy.Deploy.ID,
	)

	assert.Nil(t, deploy)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusForbidden))
}
