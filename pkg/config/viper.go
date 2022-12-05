package config

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func ConfigDir() string {
	home, err := os.UserHomeDir()
	pterm.Fatal.PrintOnError(err)
	return filepath.Join(home, ".config", "wbserver", "")
}

func InitConfig() {
	err := os.MkdirAll(ConfigDir(), os.ModePerm)
	pterm.PrintOnError(err)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AddConfigPath(ConfigDir())
	viper.AddConfigPath(".")
	viper.SetDefault("instance", "production")
	viper.SafeWriteConfig()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		pterm.PrintOnError(err)
	}
}
