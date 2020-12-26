package api

// APIEndpoints rappresent the API endpoints paths
type APIEndpoints struct {
	Apps       string
	DeployByID string
	Deploy     string
	Envs       string
}

// API has all the endpoints path
var API = &APIEndpoints{
	Apps:       "/apps",
	DeployByID: "/app/%s/deploy/%s",
	Deploy:     "/app/deploy",
	Envs:       "/app/%s/envs",
}
