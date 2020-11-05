package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetGitRoot(t *testing.T, stdout string) {
	path, err := GetGitRoot()
	assert.Equal(t, stdout, path)
	assert.Nil(t, err)
}

func TestGetGitRoot(t *testing.T) {
	runTestFakeCommand(t, "/home/g/go/src/stormkit-cli", 0, testGetGitRoot)
}

func testGetGitRootError(t *testing.T, stdout string) {
	path, err := GetGitRoot()
	assert.Equal(t, stdout, path)
	assert.Equal(t, stdout, err.Error())
}

func TestGetGitRootError(t *testing.T) {
	runTestFakeCommand(t, "exit status 128", 128, testGetGitRootError)
}
