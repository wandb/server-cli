package cmd

import "github.com/spf13/cobra"

var downgrade = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrades an instance to a previous version.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(downgrade)
}
