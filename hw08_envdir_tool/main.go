package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	envDirPath := args[0]
	cmd := args[1:]

	env, err := ReadDir(envDirPath)

	if err != nil {
		fmt.Print(err)
	}

	exitCode := RunCmd(cmd, env)
	os.Exit(exitCode)
}
