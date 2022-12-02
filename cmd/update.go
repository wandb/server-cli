package cmd

import "github.com/spf13/cobra"

var update = &cobra.Command{
	Use:   "update",
	Short: "Update an instance, configuration and license.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(update)
}
