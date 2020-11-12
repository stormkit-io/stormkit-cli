package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stormkit-io/stormkit-cli/stormkit"
)

const deployByIDapi = "/app/%s/deploy/%s"

// Deploy of application
type Deploy struct {
	ID                string `json:"id"`
	AppID             string `json:"appId"`
	Branch            string `json:"branch"`
	NumberOfFiles     string `json:"numberOfFiles"`
	Version           string `json:"version"`
	Exit              int    `json:"exit"`
	Percentage        int    `json:"percentage"`
	PullRequestNumber int    `json:"pullRequestNumber"`
	IsAutoDeploy      bool   `json:"isAutoDeploy"`
	CreatedAt         int64  `json:"createdAt"`
	StoppedAt         int64  `json:"stoppedAt"`
	IsRunning         bool   `json:"isRunning"`
	Preview           string `json:"preview"`
	Logs              []Log  `json:"logs"`
}

type SingleDeploy struct {
	Deploy Deploy `json:"deploy"`
}

func DeployByID(appID, id string) (*SingleDeploy, error) {
	// build api string
	s := fmt.Sprintf(deployByIDapi, appID, id)
	// get stormkit http client and build request
	c := stormkit.GetClient()
	request, err := stormkit.Get(s)
	if err != nil {
		return nil, err
	}

	// run request
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	// check response code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error while doing request (response: %s)", response.Status)
	}

	defer response.Body.Close()

	// convert response in SingleDeploy struct
	body, err := ioutil.ReadAll(response.Body)

	var d SingleDeploy
	err = json.Unmarshal(body, &d)

	return &d, err
}
