package api

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

// DumpLog format the data of the Log struct in a string
func DumpLog(l Log) string {
	if l.Status {
		return fmt.Sprintf("%s    OK\n\n%s", l.Title, l.Message)
	}

	return fmt.Sprintf("%s\n\n%s", l.Title, l.Message)
}
