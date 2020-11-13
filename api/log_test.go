package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedLog = Log{
	Message: "message of log",
	Status:  false,
	Title:   "title of log",
}

func TestDumpLog(t *testing.T) {
	s := DumpLog(expectedLog)

	assert.Equal(t, "title of log\n\nmessage of log", s)
}
