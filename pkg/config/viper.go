package config

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func ConfigDir() string {
	home, err := os.UserHomeDir()
	pterm.Fatal.PrintOnError(err)
	return filepath.Join(home, ".config", "wbserver", "")
}
