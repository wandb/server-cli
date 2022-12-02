package cmd

import "github.com/spf13/cobra"

var verify = &cobra.Command{
	Use:   "verify",
	Short: "Runs a few tests to make sure the instance is configured correctly.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(verify)
}
