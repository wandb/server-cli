package cmd

import "github.com/spf13/cobra"
import "fmt"

var update = &cobra.Command{
	Use:   "update",
	Short: "Update an instance, configuration and license.",
	/*
	- Updates license if a new one is generated (renewal, upsell etc.)
	- Updates wandb version
	*/
}

func license_update(env, license_key) {

}

func version_update(env) {

}

func init() {
	RootCmd.AddCommand(update)
}
