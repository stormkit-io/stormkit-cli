package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDumpApp(t *testing.T) {
	s := MockApp.Dump()
	createdAt := time.Unix(int64(MockApp.CreatedAt), 0)
	deployedAt := time.Unix(int64(MockApp.DeployedAt), 0)
	expected := fmt.Sprintf("Repo: repo0\n  ID: 1234\n  Status: false\n  AutoDeploy: \n  DefaultEnv: \n  Endpoint: \n  DisplayName: \n  CreatedAt: %s\n  DeployedAt: %s\n\n", createdAt, deployedAt)

	assert.Equal(t, expected, s)
}
