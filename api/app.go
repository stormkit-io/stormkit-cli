package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// GetApps queries the list of apps of the user
func GetApps() (*model.Apps, error) {
	// get stormkit http client and build requests
	c := stormkit.GetClient()
	request, err := stormkit.Get(API.Apps)
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

	var a model.Apps
	err = json.Unmarshal(body, &a)

	return &a, err
}
