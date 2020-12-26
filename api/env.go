package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// Envs gives an array of envs of the app Given as argument
func Envs(appID string) (*model.EnvsArray, error) {
	// get stormkit http client and build request
	c := stormkit.GetClient()
	api := fmt.Sprintf(API.Envs, appID)
	request, err := stormkit.Get(api)
	if err != nil {
		return nil, err
	}

	// do request via stormkit http client
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	// check response code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error while doing request (response: %s)", response.Status)
	}

	defer response.Body.Close()

	// convert response in EnvsArray struct
	body, err := ioutil.ReadAll(response.Body)

	var a model.EnvsArray
	err = json.Unmarshal(body, &a)

	return &a, err
}
