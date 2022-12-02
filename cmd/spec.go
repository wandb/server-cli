package cmd

import "github.com/spf13/cobra"

var spec = &cobra.Command{
	Use:   "spec",
	Short: "Returns information about instance configuration.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(spec)
}
