package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DeploymentType string

const (
	ManagedDedicatedCloud DeploymentType = "W&B Managed Dedicated Cloud"
	ManagedPrivateCloud   DeploymentType = "W&B Managed Private Cloud"
	PrivateCloud          DeploymentType = "Self-Managed Private Cloud"
	BareMetal             DeploymentType = "Self-Managed Bare Metal"
)

type DeploymentPlatform string

const (
	AWS        DeploymentPlatform = "Amazon Web Services"
	GCP        DeploymentPlatform = "Google Cloud"
	Azure      DeploymentPlatform = "Azure"
	Host       DeploymentPlatform = "Host"
	Kubernetes DeploymentPlatform = "Kubernetes"
)

type DeploymentEngine string

const (
	BYOB      DeploymentEngine = "BYOB"
	HelmChart DeploymentEngine = "Helm Chart"
	Terraform DeploymentEngine = "Terraform"
	Docker    DeploymentEngine = "Docker"
)

func CloudAuthFlow() {
	pterm.DefaultParagraph.Println(
		"You will need to sign in to your Weights & Biases Account so gain access to your license keys.",
	)
}

func GetDeploymentFlow() (string, string, string) {
	// This will change when we support context
	var deployment = "deployment"

	dtype, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select deployment type").
		WithOptions([]string{
			string(ManagedDedicatedCloud),
			string(ManagedPrivateCloud),
			string(PrivateCloud),
			string(BareMetal),
		}).
		Show()
	viper.Set(deployment+".type", dtype)

	platformOptions := []string{}
	switch dtype {
	case string(ManagedDedicatedCloud):
		fallthrough
	case string(ManagedPrivateCloud):
		fallthrough
	case string(PrivateCloud):
		platformOptions = append(platformOptions, string(AWS), string(GCP), string(Azure))
	case string(BareMetal):
		platformOptions = append(platformOptions, string(Host), string(Kubernetes))
	}
	platform, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select deployment platform").
		WithOptions(platformOptions).
		Show()
	viper.Set(deployment+".platform", platform)

	engine := string(Terraform)
	if platform == string(Host) {
		engine = string(Docker)
	}
	if platform == string(Kubernetes) {
		engine = string(HelmChart)
	}
	viper.Set(deployment+".engine", engine)
	viper.WriteConfig()
	return dtype, platform, engine
}

func ConfigureTerraformFlow() {

}

var setup = &cobra.Command{
	Use:   "setup",
	Short: "Configures and setups a W&B Server",
	Run: func(cmd *cobra.Command, args []string) {
		dtype, platform, engine := GetDeploymentFlow()
		pterm.Info.Print(dtype + " > " + platform + " > " + engine)

		if engine == string(Terraform) {
			ConfigureTerraformFlow()
		}
	},
}

func init() {
	RootCmd.AddCommand(setup)
}
