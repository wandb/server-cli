package cmd

import "os/exec"

func IsTerraformInstalled() bool {
	cmd := exec.Command("terraform", "-version")
	err := cmd.Run()
	return err == nil
}

func IsAWSCLIInstalled() bool {
	cmd := exec.Command("aws", "--version")
	err := cmd.Run()
	return err == nil
}
