package cmd

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/wandb/server-cli/pkg/auth"
	"github.com/wandb/server-cli/pkg/deployments"
	"github.com/wandb/server-cli/pkg/deployments/terraform/byob"
)

func ConfigureFlow() {
	i := deployments.GetInstance()
	dtype := i.GetType()
	engine := i.GetEngine()

	useTerraform := engine == deployments.Terraform
	isManagedDedicatedCloud := dtype == deployments.ManagedDedicatedCloud
	// isManagedPrivateCloud := dtype == deployments.ManagedPrivateCloud
	// isPrivateCloud := dtype == deployments.PrivateCloud
	// isBareMetal := dtype == deployments.BareMetal

	isBYOB := useTerraform && isManagedDedicatedCloud
	if isBYOB {
		pterm.DefaultSection.Println("Configure your bucket")
		byob.ConfigureBYOB()
		return
	}

	pterm.Fatal.Println(
		"Sorry, we currently do not support this configuration type. " +
			"Please contact support if you this this is an issue.",
	)
}

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

		confirmed, _ := pterm.DefaultInteractiveConfirm.
			WithDefaultValue(true).
			Show("Would you like to continue")
		if !confirmed {
			os.Exit(1)
		}

		auth.CloudAuthFlow()
		deployments.GetDeploymentStrategy()
		deployments.Licensing()

		ConfigureFlow()
	},
}

func init() {
	RootCmd.AddCommand(setup)
}
