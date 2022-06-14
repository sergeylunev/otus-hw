package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdName := cmd[0]
	args := cmd[1:]
	envs := envsToString(env)

	return runner(cmdName, args, envs)
}

func runner(cmd string, args []string, env []string) int {
	cmdExecutor := exec.Command(cmd, args...)

	cmdExecutor.Env = append(os.Environ(), env...)
	cmdExecutor.Stdout = os.Stdout
	cmdExecutor.Stdin = os.Stdin
	cmdExecutor.Stderr = os.Stderr

	if err := cmdExecutor.Run(); err != nil {
		return cmdExecutor.ProcessState.ExitCode()
	}

	return 0
}

func envsToString(env Environment) []string {
	envs := []string{}
	for k, v := range env {
		_, ok := os.LookupEnv(k)
		if ok {
			os.Unsetenv(k)
		}
		if !v.NeedRemove {
			envs = append(envs, k+"="+v.Value)
		}
	}

	return envs
}
