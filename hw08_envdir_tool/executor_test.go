package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("run command", func(t *testing.T) {
		cmd := []string{
			"echo",
			"test",
		}
		env := Environment{}
		result := RunCmd(cmd, env)
		require.Equal(t, 0, result)
	})
	t.Run("run command with env", func(t *testing.T) {
		cmd := []string{
			"echo",
			"test",
		}
		env := Environment{
			"test":  {"test", false},
			"empty": {"", true},
		}
		result := RunCmd(cmd, env)
		require.Equal(t, 0, result)
	})
	t.Run("error with wrong command", func(t *testing.T) {
		cmd := []string{
			"echoasfa",
			"test",
		}
		env := Environment{}
		result := RunCmd(cmd, env)
		require.Equal(t, -1, result)
	})
}
