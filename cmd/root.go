package cmd

import "github.com/spf13/cobra"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "wbserver",
	Short: "CLI tooling for W&B Server",
	Long:  `CLI tooling to help manage and deploy W&B Server`,
}
