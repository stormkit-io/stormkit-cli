package stormkit

import (
	"fmt"
	"net/http"
	"io"
)

// authorizationHeaderString
const authorizationHeaderString = "Authorization"

// client is the client for http requests
var client *http.Client

// GetClient build the http client
func GetClient() (*http.Client) {
	if client == nil {
		client = &http.Client{
			Timeout: clientTimeout,
		}
	}

	return client
}

func request(m, api string, body io.Reader) (*http.Request, error) {
	var protocol string
	if useHTTPS {
		protocol = "https"
	} else {
		protocol = "http"
	}
	protocol += "://"

	url := protocol + server + api 
	r, err := http.NewRequest(m, url, body)

	if err != nil {
		return nil, err
	}

	r.Header.Set(authorizationHeaderString, fmt.Sprintf("Bearer %s", bearerToken))

	return r, nil
}
