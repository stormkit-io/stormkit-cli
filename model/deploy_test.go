package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogDifferenceError(t *testing.T) {
	d := Deploy{ID: "abc"}
	newD := Deploy{ID: "123"}

	s, err := d.LogDifference(&newD)

	assert.Empty(t, s)
	assert.Equal(t, "cannot compare different deploys", err.Error())
}

func TestLogDifferenceEquals(t *testing.T) {
	d := Deploy{
		AppID: "123",
		Logs:  []Log{{Message: "helo"}},
	}
	newD := d

	s, err := d.LogDifference(&newD)

	assert.Equal(t, "", s)
	assert.Nil(t, err)
}

func TestLogDifference(t *testing.T) {
	d := Deploy{
		AppID: "123",
		Logs:  []Log{{Message: "helo"}},
	}
	newD := d
	newD.Logs = []Log{
		{Message: "helo one"},
		{Message: "helo two"},
	}

	s, err := d.LogDifference(&newD)

	assert.Equal(t, " one\nhelo two\n", s)
	assert.Nil(t, err)
}
