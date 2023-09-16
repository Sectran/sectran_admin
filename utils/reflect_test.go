package utils

import (
	"sectran/user/config"
	"testing"

	"gotest.tools/assert"
)

func TestReflect(t *testing.T) {
	config := config.SSHConfig{}
	config.UserName = "111"

	SetVal(&config, "UserName", "ryan")
	assert.Equal(t, config.UserName, "ryan")
}
