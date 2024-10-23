package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersionCommand(t *testing.T) {
	cmd := NewVersionCommand()
	cmd.SetArgs([]string{
		"version",
	})
	err := cmd.Execute()
	assert.Nil(t, err)
}
