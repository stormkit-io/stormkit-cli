package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stormkit-io/stormkit-cli/stormkit"
)

const dumpAppPrintf = "Repo: %s\n  ID: %s\n  Status: %t\n  AutoDeploy: %s\n  DefaultEnv: %s\n  Endpoint: %s\n  DisplayName: %s\n  CreatedAt: %s\n  DeployedAt: %s\n\n"
const getApps = "/apps"

// App is the model of an app
type App struct {
	// ID of the app
	ID string `json:"id"`
	// Repo is the repository of the app
	Repo string `json:"repo"`
	// DisplayName UNKOWN
	DisplayName string `json:"displyName"`
	// CreatedAt is the UNIX time of the creatino of the app
	CreatedAt int `json:"createdAt"`
	// DefaultEnv is the default environment
	DefaultEnv string `json:"defaultEnv"`
	// DeployedAt is the UNIX time of the last deployment
	DeployedAt int `json:"deployedAt"`
	// Status of the last deployment
	Status bool `json:"status"`
	// UserID of the user signed in
	UserID string `json:"userId"`
	// Endpoint is the endpoint doamin (not shure)
	Endpoint string `json:"endpoint"`
	// AutoDeploy type (commit/pr)
	AutoDeploy string `json:"autoDeploy"`
}

// Apps object conatining list of apps
type Apps struct {
	// Apps list of the apps
	Apps []App `json:"apps"`
}

// GetApps queries the list of apps of the user
func GetApps() (*Apps, error) {
	// get stormkit http client and build requests
	c := stormkit.GetClient()
	request, err := stormkit.Get(getApps)
	if err != nil {
		return nil, err
	}

	// run requests
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	// check response code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error while doing request (response: %s)", response.Status)
	}

	defer response.Body.Close()

	// convert response in Apps object
	body, err := ioutil.ReadAll(response.Body)

	var a Apps
	err = json.Unmarshal(body, &a)

	return &a, err
}

// DumpApp in a string with all the parameters of the app
func DumpApp(a App) string {
	return fmt.Sprintf(
		dumpAppPrintf,
		a.Repo,
		a.ID,
		a.Status,
		a.AutoDeploy,
		a.DefaultEnv,
		a.Endpoint,
		a.DisplayName,
		time.Unix(int64(a.CreatedAt), 0),
		time.Unix(int64(a.DeployedAt), 0),
	)
}
