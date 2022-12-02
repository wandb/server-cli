package cmd

import "github.com/spf13/cobra"

var spec = &cobra.Command{
	Use:   "spec",
	Short: "Returns information about instance configuration.",
}

func init() {
	RootCmd.AddCommand(spec)
}
