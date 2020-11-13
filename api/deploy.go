package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// DeployByIDapi is the api string formatter for Sprintf(), first argument
// app id, second argument deploy id
const DeployByIDapi = "/app/%s/deploy/%s"

// DeployByID calls the stormkit http api with the appID and the log id,
// it returns a SingleDeploy struct that is the rappresentation of the
// http response
func DeployByID(appID, id string) (*model.SingleDeploy, error) {
	// build api string
	s := fmt.Sprintf(DeployByIDapi, appID, id)
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
