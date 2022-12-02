package cmd

import "github.com/spf13/cobra"

var logs = &cobra.Command{
	Use:   "logs",
	Short: "Downloads instance debug bundle.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(logs)
}
