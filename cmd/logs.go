package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	// "github.com/wandb/server-cli/pkg/deployments"
)

/* 
   Ideally debug-bundle should be available to download
   from an endpoint behind some form of authentication. 
   We could ask user to enter their API key to determine
   if they're admin or not. Then either allow or deny the
   request for the debug bundle.
*/

var logs = &cobra.Command{
	Use:   "logs",
	Short: "Downloads instance debug bundle.",
	Run: func(cmd *cobra.Command, args []string) {
		// i := deployments.GetInstance()
		pterm.Println("TODO:Define how it should download debug bundle")
	},
}

func init() {
	RootCmd.AddCommand(logs)
}
