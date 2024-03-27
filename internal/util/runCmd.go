package util

import (
	"os/exec"
	"strings"
)

func RunCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	return cmd.Run()
}

func RunCommandWithOutput(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
