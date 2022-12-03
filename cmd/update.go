package cmd

import "github.com/spf13/cobra"

var update = &cobra.Command{
	Use:   "update",
	Short: "Update an instance, configuration and license.",
}

func init() {
	RootCmd.AddCommand(update)
}
