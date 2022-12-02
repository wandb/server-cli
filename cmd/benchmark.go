package cmd

import "github.com/spf13/cobra"

var benchmark = &cobra.Command{
	Use:   "benchmark",
	Short: "Run a benchmark to determine performance",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	RootCmd.AddCommand(benchmark)
}
