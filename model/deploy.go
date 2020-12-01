package model

import (
	"fmt"
)

// Deploy data of application
type Deploy struct {
	ID                string `json:"id"`
	AppID             string `json:"appId"`
	Branch            string `json:"branch"`
	NumberOfFiles     int    `json:"numberOfFiles"`
	Version           string `json:"version"`
	Exit              int    `json:"exit"`
	Env               string `json:"env"`
	Percentage        int    `json:"percentage"`
	PullRequestNumber int    `json:"pullRequestNumber"`
	IsAutoDeploy      bool   `json:"isAutoDeploy"`
	CreatedAt         int64  `json:"createdAt"`
	StoppedAt         int64  `json:"stoppedAt"`
	IsRunning         bool   `json:"isRunning"`
	Preview           string `json:"preview"`
	Logs              []Log  `json:"logs"`
}

// LastLog returns the latest log of the deploy
func (d *Deploy) LastLog() *Log {
	if len(d.Logs) == 0 {
		return nil
	}

	return &d.Logs[len(d.Logs)-1]
}

// LogDifference retuns the difference of the logs with the one given as arguemnt
func (d *Deploy) LogDifference(newD *Deploy) (string, error) {
	// cannot compare differents deploys logs
	if d.ID != "" && d.ID != newD.ID {
		return "", fmt.Errorf("cannot compare different deploys")
	}

	// if no logs in new deploy return empty
	if len(newD.Logs) == 0 {
		return "", nil
	}

	var oldLogMessage, newLogMessage, s string

	// read last log from old deploy
	if l := d.LastLog(); l != nil {
		oldLogMessage = l.Message
	}
	// read latest readed log from new deploy (via old deploy log)
	if len(d.Logs) == 0 {
		//		newLogMessage = newD.Logs[0].Message
	} else {
		newLogMessage = newD.Logs[len(d.Logs)-1].Message
	}

	// append new pieces of latest readed log
	if len(newLogMessage) > len(oldLogMessage) {
		s += newLogMessage[len(oldLogMessage):]
	}

	// append all new logs
	if len(newD.Logs) > len(d.Logs) {
		for i := len(d.Logs); i < len(newD.Logs); i++ {
			s += "\n" + newD.Logs[i].Message
		}
		s += "\n"
	}

	return s, nil
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
