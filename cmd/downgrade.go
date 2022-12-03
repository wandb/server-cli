package cmd

import "github.com/spf13/cobra"

var downgrade = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrades an instance to a previous version.",
}

func init() {
	RootCmd.AddCommand(downgrade)
}
