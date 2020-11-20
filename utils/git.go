package utils

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

// GitRoot gives the root of the git repository where executed
func GitRoot() (string, error) {
	pathBytes, err := execCommand("git", "rev-parse", "--show-toplevel").Output()
	path := strings.TrimSpace(string(pathBytes))

	return path, err
}

// GitBranchesNames retrives the branches as strings
func GitBranchesNames(r *git.Repository) ([]string, error) {
	ri, err := r.Branches()
	if err != nil {
		return nil, err
	}

	var a []string

	for b, err := ri.Next(); err == nil; b, err = ri.Next() {
		a = append(a, b.Name().String())
	}

	for i, v := range a {
		a[i] = v[strings.LastIndex(v, "/")+1:]
	}

	return a, err
}
