package model

import (
	"fmt"
	"time"
)

// dumpAppPrintf is the formatting string of the app Dump
const dumpAppPrintf = "Repo: %s\n  ID: %s\n  Status: %t\n  AutoDeploy: %s\n  DefaultEnv: %s\n  Endpoint: %s\n  DisplayName: %s\n  CreatedAt: %s\n  DeployedAt: %s\n\n"

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

// MockApp mock data for App struct
var MockApp = App{
	Repo: "repo0",
	ID:   "1234",
}

// MockApps mock data for Apps struct
var MockApps = Apps{
	Apps: []App{
		MockApp,
		{
			Repo: "repo1",
			ID:   "1235",
		},
	},
}

// Dump returns the app data formatted
func (a *App) Dump() string {
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
