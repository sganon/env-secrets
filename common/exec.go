package common

import (
	"encoding/json"
	"os/exec"
)

// ExecCLI exec a shell command and tries to decode its output as JSON
func ExecCLI(binaryName string, args []string, output interface{}) error {
	cmd := exec.Command(binaryName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if err := json.NewDecoder(stdout).Decode(&output); err != nil {
		return err
	}
	return nil
}
