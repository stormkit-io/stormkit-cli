package model

import (
	"fmt"
)

// Log of an operation of a deployment
type Log struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Title   string `json:"title"`
	Payload struct {
		Branch string `json:"branch"`
		Commit struct {
			Author  string `json:"author"`
			Message string `json:"message"`
			SHA     string `json:"sha"`
		}
	}
}

// MockLogs array of istances of Log struct
var MockLogs = []Log{
	{
		Title:   "title log 0",
		Message: "message log 0 that is true",
		Status:  true,
	},
	{
		Title:   "title log 1",
		Message: "message log 1 that is false",
		Status:  false,
	},
}

// Dump format the data of the Log struct in a string
func (l *Log) Dump() string {
	if l.Status {
		return fmt.Sprintf("%s    OK\n\n%s", l.Title, l.Message)
	}

	return fmt.Sprintf("%s\n\n%s", l.Title, l.Message)
}
