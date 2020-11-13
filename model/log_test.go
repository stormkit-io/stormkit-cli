package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDumpLog(t *testing.T) {
	s := MockLogs[0].Dump()

	assert.Equal(t, "title log 0    OK\n\nmessage log 0 that is true", s)
}
