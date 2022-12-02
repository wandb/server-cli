package cmd

import "github.com/spf13/cobra"

var benchmark = &cobra.Command{
	Use:   "benchmark",
	Short: "Run a benchmark to determine performance",
}

func init() {
	RootCmd.AddCommand(benchmark)
}
