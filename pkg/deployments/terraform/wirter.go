package terraform

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/config"
	"github.com/wandb/server-cli/pkg/deployments"
)

func GetTerraformPath() {
	name := deployments.GetInstance().GetName()
	terraformPath := filepath.Join(config.ConfigDir(), name)

	err := os.MkdirAll(terraformPath, os.ModePerm)
	pterm.PrintOnError(err)
}
