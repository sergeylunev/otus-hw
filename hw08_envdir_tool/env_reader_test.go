package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read empty env file from directory", func(t *testing.T) {
		env, err := ReadDir("testdata/empty_env")
		envValue := EnvValue{
			Value:      "",
			NeedRemove: true,
		}
		expectedEnv := make(Environment)
		expectedEnv["EMPTY"] = envValue

		require.Nil(t, err)
		require.EqualValues(t, expectedEnv, env)
	})
	t.Run("read simple file from directory", func(t *testing.T) {
		env, err := ReadDir("testdata/simple_env")

		envValue := EnvValue{
			Value:      "env",
			NeedRemove: false,
		}
		expectedEnv := make(Environment)
		expectedEnv["SIMPLE"] = envValue

		require.Nil(t, err)
		require.EqualValues(t, expectedEnv, env)
	})
	t.Run("read file with two lines", func(t *testing.T) {
		env, err := ReadDir("testdata/two_lines_env")
		envValue := EnvValue{
			Value:      "env",
			NeedRemove: false,
		}
		expectedEnv := make(Environment)
		expectedEnv["TWOLINES"] = envValue

		require.Nil(t, err)
		require.EqualValues(t, expectedEnv, env)
	})
	t.Run("remove whitespace characters in the end of string", func(t *testing.T) {
		env, err := ReadDir("testdata/whitespace_env")
		envValue := EnvValue{
			Value:      "env",
			NeedRemove: false,
		}
		expectedEnv := make(Environment)
		expectedEnv["WHITESPACE"] = envValue

		require.Nil(t, err)
		require.EqualValues(t, expectedEnv, env)
	})
	t.Run("env with new line", func(t *testing.T) {
		env, err := ReadDir("testdata/with_new_line_env")
		envValue := EnvValue{
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		}
		expectedEnv := make(Environment)
		expectedEnv["FOO"] = envValue

		require.Nil(t, err)
		require.EqualValues(t, expectedEnv, env)
	})
	t.Run("no equal sign in file name", func(t *testing.T) {
		_, err := ReadDir("testdata/wrong_file_name_env")

		require.NotNil(t, err)
		require.ErrorIs(t, err, ErrFileName)
	})
}
