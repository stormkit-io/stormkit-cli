package utils

import (
	"os"
	"os/exec"
	"strconv"
	"testing"
)

// goWantHelperProcessString is string for env var GO_WANT_HELP_PROCESS (testing flag)
const goWantHelperProcessString = "GO_WANT_HELPER_PROCESS"

// stdoutString is the string for env var STDOUT (requested output in test)
const stdoutString = "STDOUT"

// exitStatusString is the exist status for the env var EXIT_STATUS (requested exit status in test)
const exitStatusString = "EXIT_STATUS"

// mockedExitStatus is the exit status that will be returned in the test
var mockedExitStatus = 0

// mockedStdout is the  stdout that will be returned in the test
var mockedStdout string

// execCommand is local var for run&teest command
var execCommand = exec.Command

// fakeExecCommand creates a fake executable command for tests internal of utils package
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecCommandHelper", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(mockedExitStatus)
	cmd.Env = []string{
		goWantHelperProcessString + "=1",
		stdoutString + "=" + mockedStdout,
		exitStatusString + "=" + es,
	}

	return cmd
}

// runTestFakeCommand execute a test in a sub process with `exec.Command` function
func runTestFakeCommand(t *testing.T, stdout string, exitStatus int, f func(t *testing.T, stdout string)) {
	mockedStdout = stdout
	mockedExitStatus = exitStatus
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	f(t, stdout)
}
