package utils

import (
	"strings"
)

// GitRoot gives the root of the git repository where executed
func GitRoot() (string, error) {
	pathBytes, err := execCommand("git", "rev-parse", "--show-toplevel").Output()
	path := strings.TrimSpace(string(pathBytes))

	return path, err
}
