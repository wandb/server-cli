package cmd

import (
	"os/exec"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/config"
)

func CloudAuthFlow() {
	pterm.DefaultParagraph.Println(
		"You will need to sign in to your Weights & Biases Account so gain access to your license keys.",
	)
	auth := "https://wandb.ai/authorize"
	pterm.Blue("You can find your API key in your browser here: https://wandb.ai/authorize")

	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", auth).Start()
	case "windows", "darwin":
		exec.Command("open", auth).Start()
	default:
		pterm.Warning.Print("Could not find default browser")
	}

	apikey, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Pate an API key from your profile and hit enter or press ctrl+c to quit").Show()
	viper.Set("wandb.apikey", apikey)
	viper.WriteConfig()
}

func GetDeploymentFlow() (string, string, string) {

	instance := config.GetInstance()
	dtype, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select deployment type").
		WithOptions([]string{
			string(config.ManagedDedicatedCloud),
			string(config.ManagedPrivateCloud),
			string(config.PrivateCloud),
			string(config.BareMetal),
		}).
		Show()

	instance.SetType(dtype)

	platformOptions := []string{}
	switch dtype {
	case string(config.ManagedDedicatedCloud):
		fallthrough
	case string(config.ManagedPrivateCloud):
		fallthrough
	case string(config.PrivateCloud):
		platformOptions = append(platformOptions, string(config.AWS), string(config.GCP), string(config.Azure))
	case string(config.BareMetal):
		platformOptions = append(platformOptions, string(config.Host), string(config.Kubernetes))
	}
	platform, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select deployment platform").
		WithOptions(platformOptions).
		Show()

	instance.SetPlatform(dtype)

	engine := string(config.Terraform)
	if platform == string(config.Host) {
		engine = string(config.Docker)
	}
	if platform == string(config.Kubernetes) {
		engine = string(config.HelmChart)
	}
	instance.SetPlatform(engine)
	instance.Write()

	return dtype, platform, engine
}

func ConfigureTerraformFlow() {

}

var setup = &cobra.Command{
	Use:   "setup",
	Short: "Configures and setups a W&B Server",
	Run: func(cmd *cobra.Command, args []string) {
		CloudAuthFlow()
		dtype, platform, engine := GetDeploymentFlow()
		pterm.Info.Println(dtype + " > " + platform + " > " + engine)
		if engine == string(config.Terraform) {
			ConfigureTerraformFlow()
		}
	},
}

func init() {
	RootCmd.AddCommand(setup)
}
