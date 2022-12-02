package cmd

import "github.com/spf13/cobra"

var logs = &cobra.Command{
	Use:   "logs",
	Short: "Downloads instance debug bundle.",
}

func init() {
	RootCmd.AddCommand(logs)
}
