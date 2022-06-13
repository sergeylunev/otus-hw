package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrFileName = errors.New("wrong env file name")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := Environment{}

	for _, f := range files {
		f.Name()
		if strings.Contains(f.Name(), "=") {
			return nil, ErrFileName
		}
		if v, _ := f.Info(); v.Size() == 0 {
			env[f.Name()] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
			continue
		}
		bytes, err := os.ReadFile(dir + "/" + f.Name())

		if err != nil {
			return nil, err
		}

		result, err := readEnvFromFile(bytes)
		if err != nil {
			return nil, err
		}
		env[f.Name()] = EnvValue{
			Value:      result,
			NeedRemove: false,
		}

		fmt.Println(result)
	}

	if err != nil {
		return nil, err
	}

	return env, nil
}

func readEnvFromFile(b []byte) (string, error) {
	result := make([]byte, 0)
	for _, ch := range b {
		if ch == 10 {
			break
		}
		if ch == 0 {
			ch = 10
		}
		result = append(result, ch)
	}
	return strings.TrimRight(string(result), "\t "), nil
}
