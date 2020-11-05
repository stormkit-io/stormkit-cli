package stormkit

import (
	"fmt"
	"io"
	"net/http"
)

// authorizationHeaderString
const authorizationHeaderString = "Authorization"

// client is the client for http requests
var client *http.Client

// GetClient build the http client
func GetClient() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout: globalConfig.ClientTimeout,
		}
	}

	return client
}

func request(m, api string, body io.Reader) (*http.Request, error) {
	var protocol string
	if globalConfig.UseHTTPS {
		protocol = "https"
	} else {
		protocol = "http"
	}
	protocol += "://"

	url := protocol + globalConfig.Server + api
	r, err := http.NewRequest(m, url, body)

	if err != nil {
		return nil, err
	}

	r.Header.Set(authorizationHeaderString, fmt.Sprintf("Bearer %s", globalConfig.BearerToken))

	return r, nil
}

// Get create a get request for stormkit
func Get(api string) (*http.Request, error) {
	return request(http.MethodGet, api, nil)
}

// Post create a post request for stormkit
func Post(a string, b io.Reader) (*http.Request, error) {
	return request(http.MethodPost, a, b)
}

// Put create a put request for stormkit
func Put(a string, b io.Reader) (*http.Request, error) {
	return request(http.MethodPut, a, b)
}

// Delete create a delete request for stormkit
func Delete(a string, b io.Reader) (*http.Request, error) {
	return request(http.MethodDelete, a, b)
}
