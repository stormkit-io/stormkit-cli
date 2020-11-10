package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGitRoot(t *testing.T, stdout string) {
	path, err := GitRoot()
	assert.Equal(t, stdout, path)
	assert.Nil(t, err)
}

func TestGitRoot(t *testing.T) {
	runTestFakeCommand(t, "/home/g/go/src/stormkit-cli", 0, testGitRoot)
}

func testGitRootError(t *testing.T, stdout string) {
	path, err := GitRoot()
	assert.Equal(t, stdout, path)
	assert.Equal(t, stdout, err.Error())
}

func TestGetGitRootError(t *testing.T) {
	runTestFakeCommand(t, "exit status 128", 128, testGitRootError)
}
