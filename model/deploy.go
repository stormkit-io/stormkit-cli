package model

// Deploy data of application
type Deploy struct {
	ID                string `json:"id"`
	AppID             string `json:"appId"`
	Branch            string `json:"branch"`
	NumberOfFiles     int    `json:"numberOfFiles"`
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

// SingleDeploy is struct containing a Deploy
type SingleDeploy struct {
	Deploy Deploy `json:"deploy"`
}

// MockDeploy mock istance of Deploy struct
var MockDeploy = Deploy{
	ID:    "12345",
	AppID: "12346",
	Logs:  MockLogs,
}

// MockSingleDeploy is an istance of Single Deploy
var MockSingleDeploy = SingleDeploy{
	Deploy: MockDeploy,
}
