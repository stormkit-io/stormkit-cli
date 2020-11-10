package utils

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

// TestExecCommandHelper is the fake test to print the right stdout and exit status
func TestExecCommandHelper(t *testing.T) {
	if os.Getenv(goWantHelperProcessString) != "1" {
		return
	}

	fmt.Fprintf(os.Stdout, os.Getenv(stdoutString))
	i, _ := strconv.Atoi(os.Getenv(exitStatusString))
	os.Exit(i)
}
