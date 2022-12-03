package cmd

import "github.com/spf13/cobra"

var spec = &cobra.Command{
	Use:   "spec",
	Short: "Returns information about instance configuration.",
<<<<<<< HEAD
	/*
	- Deployment Region
	- MySQL db type, version and instance type
	- Redis size and memory?
	- K8s node/EC2 instance size and type
	*/
=======
>>>>>>> main
}

func init() {
	RootCmd.AddCommand(spec)
}
