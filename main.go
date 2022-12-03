package main

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/cmd"
)

func configureTheme() {
	pterm.EnableDebugMessages()
	pterm.EnableColor()
	pterm.DefaultInteractiveSelect.MaxHeight = 15
	pterm.Debug.Prefix.Text = "🐞"
	pterm.Debug.Prefix.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Info.Prefix.Text = "ⓘ"
	pterm.Info.Prefix.Style = &pterm.ThemeDefault.InfoMessageStyle
	pterm.Info.MessageStyle = &pterm.ThemeDefault.DefaultText
	pterm.Success.Prefix.Text = "✓"
	pterm.Success.Prefix.Style = &pterm.ThemeDefault.SuccessMessageStyle
	pterm.Warning.Prefix.Text = "!"
	pterm.Warning.Prefix.Style = &pterm.ThemeDefault.WarningMessageStyle
	pterm.Error.Prefix.Text = "✗"
	pterm.Error.Prefix.Style = &pterm.ThemeDefault.ErrorMessageStyle
	pterm.Fatal.Prefix.Text = "🤯"
	pterm.Fatal.Prefix.Style = &pterm.ThemeDefault.FatalMessageStyle
}

func initConfig() {
	home, err := os.UserHomeDir()
	pterm.Fatal.PrintOnError(err)

	configDir := filepath.Join(home, ".config", "wbserver", "")
	err = os.MkdirAll(configDir, os.ModePerm)
	pterm.PrintOnError(err)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")
	viper.SetDefault("context", "prod")
	viper.SafeWriteConfig()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		pterm.PrintOnError(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	configureTheme()
}

func main() {
	cmd.RootCmd.Execute()
	viper.WriteConfig()
}
