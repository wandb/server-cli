package cmd

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/wandb/server-cli/pkg/auth"
	"github.com/wandb/server-cli/pkg/deployments"
)

var setup = &cobra.Command{
	Use:   "setup",
	Short: "Configures and setups a W&B Server",
	Run: func(cmd *cobra.Command, args []string) {

		pterm.Println()
		pterm.DefaultParagraph.Println(
			pterm.Yellow(
				"We recommend that you consider using the https://wandb.ai cloud before privately " +
					"hosting a W&B Server on your infrastructure. The cloud is simple and secure, " +
					"with no configuration required.",
			),
		)
		pterm.Println()
		pterm.DefaultParagraph.Println(
			"Now we will walk you though setting up a W&B Server. If you exit at anytime we " +
				"will save the state and continue from where you left off.",
		)

		// pterm.Println()
		// pterm.Println("1. Cloud Authentication")
		// pterm.Println("2. Deployment Strategy")
		// pterm.Println("3. Licensing")
		// pterm.Println("4. Deploy Configuration")
		// pterm.Println("5. Instance Authentication (if required)")
		// pterm.Println("5. Instance Testing (if required)")
		// pterm.Println()

		confirmed, _ := pterm.DefaultInteractiveConfirm.
			WithDefaultValue(true).
			Show("Would you like to continue")
		if !confirmed {
			os.Exit(1)
		}

		auth.CloudAuthFlow()
		deployments.GetDeploymentStrategy()
		deployments.Licensing()

		i := deployments.GetInstance()
		dtype := i.GetType()
		// platform := i.GetPlatform()
		engine := i.GetEngine()

		useTerraform := engine == deployments.Terraform
		isManagedDedicatedCloud := dtype == deployments.ManagedDedicatedCloud
		isBYOB := useTerraform && isManagedDedicatedCloud
		if isBYOB {
			pterm.Println("Configure BYOB")
		}

		if isBYOB {
			pterm.Println("Configure BYOb")
		}

		if isBYOB {
			pterm.Println("Configure BYOb")
		}

		pterm.Error.Println(
			"Sorry, we currently do not support this configuration type. " +
				"Please contact support if you this this is an issue.",
		)
	},
}

func init() {
	RootCmd.AddCommand(setup)
}
