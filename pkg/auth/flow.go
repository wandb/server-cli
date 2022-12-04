package auth

import (
	"os/exec"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func CloudAuthFlow() {
	if viper.GetString("wandb.apikey") != "" {
		return
	}

	pterm.DefaultSection.Println("Authenticate")

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
