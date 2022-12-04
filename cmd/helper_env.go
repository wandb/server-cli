package cmd

import "os/exec"

func IsAWSCLIInstalled() bool {
	cmd := exec.Command("aws", "--version")
	err := cmd.Run()
	return err == nil
}
