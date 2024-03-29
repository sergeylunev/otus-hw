package main

import (
	"errors"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

const newLine byte = 10
const terminalByte byte = 0

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
		bytes, err := os.ReadFile(path.Join(dir, f.Name()))

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
	}

	if err != nil {
		return nil, err
	}

	return env, nil
}

func readEnvFromFile(b []byte) (string, error) {
	result := make([]byte, 0, len(b))
	for _, ch := range b {
		if ch == newLine {
			break
		}
		if ch == terminalByte {
			ch = newLine
		}
		result = append(result, ch)
	}
	return strings.TrimRight(string(result), "\t "), nil
}
