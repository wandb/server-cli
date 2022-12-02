package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configures and setups a W&B Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello called")
	},
}

func init() {
	RootCmd.AddCommand()
}
