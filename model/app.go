package model

// App is the model of an App deployed on Stormkit
type App struct {
	Repo string `json:"repo"`
	DisplayName string `json:"displyName"`
	CreatedAt int `json:"createdAt"`
	DefaultEnv string `json:"defaultEnv"`
	DeployedAt int `json:"deployedAt"`
	Status bool `json:"status"`
	UserID string `json:"userId"`
	Endpoint string `json:"endpoint"`
	AutoDeploy string `json:"autoDeploy"`
}


