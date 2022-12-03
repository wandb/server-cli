package cmd

import "github.com/spf13/cobra"

var spec = &cobra.Command{
	Use:   "spec",
	Short: "Returns information about instance configuration.",
	/*
	- Deployment Region
	- MySQL db type, version and instance type
	- Redis size and memory?
	- K8s node/EC2 instance size and type
	*/
}

func init() {
	RootCmd.AddCommand(spec)
}
