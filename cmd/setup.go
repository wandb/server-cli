package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/wandb/server-cli/pkg/auth"
	"github.com/wandb/server-cli/pkg/deployments"
)

func GetDeploymentFlow() (string, string, string) {
	instance := deployments.GetInstance()

	dtype := instance.GetType()
	if dtype == "" {
		dtype, _ = pterm.DefaultInteractiveSelect.
			WithDefaultText("Select deployment type").
			WithOptions([]string{
				string(deployments.ManagedDedicatedCloud),
				string(deployments.ManagedPrivateCloud),
				string(deployments.PrivateCloud),
				string(deployments.BareMetal),
			}).
			Show()

		instance.SetType(dtype)
	}

	platform := instance.GetPlatform()
	if platform == "" {
		platformOptions := []string{}
		switch dtype {
		case string(deployments.ManagedDedicatedCloud):
			fallthrough
		case string(deployments.ManagedPrivateCloud):
			fallthrough
		case string(deployments.PrivateCloud):
			platformOptions = append(platformOptions, string(deployments.AWS), string(deployments.GCP), string(deployments.Azure))
		case string(deployments.BareMetal):
			platformOptions = append(platformOptions, string(deployments.Host), string(deployments.Kubernetes))
		}
		platform, _ = pterm.DefaultInteractiveSelect.
			WithDefaultText("Select deployment platform").
			WithOptions(platformOptions).
			Show()

		instance.SetPlatform(platform)
	}

	engine := instance.GetEngine()
	if engine == "" {
		engine = string(deployments.Terraform)
		if platform == string(deployments.Host) {
			engine = string(deployments.Docker)
		}
		if platform == string(deployments.Kubernetes) {
			engine = string(deployments.HelmChart)
		}
		instance.SetEngine(engine)
	}
	instance.Write()

	return dtype, platform, engine
}

var setup = &cobra.Command{
	Use:   "setup",
	Short: "Configures and setups a W&B Server",
	Run: func(cmd *cobra.Command, args []string) {
		auth.CloudAuthFlow()
		deployments.CreateDeployment()

		dtype, platform, engine := GetDeploymentFlow()
		pterm.Bold.Println("Select Deployment Stratgy")
		pterm.Println(
			pterm.Green(dtype) +
				" > " + pterm.Green(platform) + " > " +
				pterm.Green(engine),
		)

		pterm.Println()

		// flow.ConfigureTerraformFlow(platform)
		// if engine == string(config.Terraform) {
		// 	ConfigureTerraformFlow()
		// }
	},
}

func init() {
	RootCmd.AddCommand(setup)
}
