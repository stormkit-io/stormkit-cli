package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// DeployByID calls the stormkit http api with the appID and the log id,
// it returns a SingleDeploy struct that is the rappresentation of the
// http response
func DeployByID(appID, id string) (*model.SingleDeploy, error) {
	// build api string
	s := fmt.Sprintf(API.DeployByID, appID, id)
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

	var d model.SingleDeploy
	err = json.Unmarshal(body, &d)

	return &d, err
}

// Deploy calls the stormkit http api for start a deploy of a determinate
// branch in an environment of an application
func Deploy(d model.Deploy) (*model.Deploy, error) {

	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	// get stormkit http client add build request
	c := stormkit.GetClient()
	request, err := stormkit.Post(API.Deploy, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error while doing request (response: %s)", response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &d)

	return &d, err
}

// WatchDeploy print the logs of the deploy
func WatchDeploy(deployID string) error {
	var oldDeploy model.Deploy

	for {
		d, err := DeployByID(stormkit.GetEngineAppID(), deployID)
		if err != nil {
			return err
		}

		log, err := oldDeploy.LogDifference(&d.Deploy)
		if err != nil {
			return err
		}

		fmt.Print(log)

		oldDeploy = d.Deploy

		if !d.Deploy.IsRunning {
			break
		}
		time.Sleep(time.Second * 2)
	}

	if l := oldDeploy.LastLog(); l != nil {
		if l.Status {
			fmt.Println("\n\nBuild Success!")
			return nil
		} else {
			return fmt.Errorf("\n\nerror while building!")
		}
	}

	return fmt.Errorf("\n\nno logs available")
}
