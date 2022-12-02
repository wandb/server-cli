package cmd

import "github.com/spf13/cobra"

var verify = &cobra.Command{
	Use:   "verify",
	Short: "Runs a few tests to make sure the instance is configured correctly.",
}

func init() {
	RootCmd.AddCommand(verify)
}
