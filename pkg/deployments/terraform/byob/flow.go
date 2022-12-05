package byob

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/deployments"
)

func ConfigureBYOB() {
	i := deployments.GetInstance()
	platform := i.GetPlatform()

	canDoByob := i.GetEngine() == deployments.Terraform &&
		i.GetType() == deployments.ManagedDedicatedCloud
	if !canDoByob {
		pterm.Fatal.Println("Invalid BYOB configuration.")
	}

	textInput := pterm.DefaultInteractiveTextInput

	cfg := new(BYOBConfig)

	bucketPrefix, _ := textInput.Show("Bucket prefix")
	bucketPrefix = strings.ToLower(bucketPrefix)
	bucketPrefix = strings.ReplaceAll(bucketPrefix, " ", "-")
	cfg.BucketPrefix = bucketPrefix

	supportedRegions := deployments.SupportedRegions(platform)
	cfg.Region, _ = pterm.
		DefaultInteractiveSelect.
		WithOptions(supportedRegions).
		Show("Select region")

	if platform == deployments.GCP {
		cfg.Google = new(GoogleBYOBConfig)
		cfg.Google.ProjectID, _ = textInput.Show("Project ID")
	}

	i.SetInterface("byob", cfg)
	i.Write()
}
