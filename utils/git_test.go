package utils

import (
	"testing"

	"github.com/go-git/go-git/v5"
	//	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
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

func TestGitBranchesNames(t *testing.T) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/go-git/go-git",
	})
	assert.Nil(t, err)

	bf, err := r.Branches()
	counter := 0
	bf.ForEach(func(v *plumbing.Reference) error {
		counter++
		return nil
	})

	branches, err := GitBranchesNames(r)

	assert.Equal(t, counter, len(branches))
	assert.Nil(t, err)
}
