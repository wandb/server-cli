package auth

import (
	"os/exec"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/api/wandb"
)

func CloudAuthFlow() {
	pterm.DefaultSection.Println("Cloud Authentication")

	if viper.GetString("wandb.apikey") != "" {
		viewer, err := wandb.GetViewer()
		pterm.Fatal.PrintOnError(err)
		pterm.Success.Print("You are signed in as, ")
		pterm.Bold.Print(pterm.Green(viewer.Name))
		pterm.Println()
		return
	}

	pterm.DefaultParagraph.Println(
		"You will need to sign in to your Weights & Biases Cloud Account " +
			"for managing and accessing license keys. Please use your account " +
			"that is linked to your work email.",
	)
	pterm.Println()
	auth := "https://wandb.ai/authorize"
	pterm.DefaultParagraph.Println("You can access your API key here:")
	pterm.Bold.Println(pterm.Cyan("\thttps://wandb.ai/authorize"))
	pterm.Println()

	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", auth).Start()
	case "windows", "darwin":
		exec.Command("open", auth).Start()
	}

	apikey, _ := pterm.DefaultInteractiveTextInput.
		Show("Paste your API key")

	viper.Set("wandb.apikey", apikey)

	viewer, err := wandb.GetViewer()
	pterm.Fatal.PrintOnError(err)

	pterm.Println()
	pterm.Print("Hello, ")
	pterm.Bold.Println(pterm.Cyan(viewer.Name))
	pterm.Println()

	viper.WriteConfig()
}
