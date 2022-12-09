package privatecloud

import (
	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/deployments"
)

type DatabaseConfig struct {
	Version string
	Size    string
}

func ConfigureDatabase(platform deployments.DeploymentPlatform) *DatabaseConfig {
	config := new(DatabaseConfig)

	var err error

	sizes := RecommendDatabaseSizes(platform)
	config.Size, err = pterm.DefaultInteractiveSelect.WithOptions(sizes).Show("Select database size")
	pterm.Fatal.PrintOnError(err)

	config.Version = GetDatabaseVersion(platform)

	return config
}

func GetDatabaseVersion(platform deployments.DeploymentPlatform) string {
	switch platform {
	case deployments.AWS:
		return "8.0.mysql_aurora.3.02.0"
	case deployments.Azure:
		return "8.0.21"
	case deployments.GCP:
		return "MYSQL_8"
	}

	pterm.Fatal.Printfln("'%s' does not support database versioning", platform)
	return ""
}

func RecommendDatabaseSizes(platform deployments.DeploymentPlatform) []string {
	switch platform {
	case deployments.AWS:
		return []string{
			"db.r5.large",
			"db.r5.xlarge",
			"db.r5.2xlarge",
			"db.r5.4xlarge",
			"db.r5.8xlarge",
			"db.r5.16xlarge",
			"db.r5.24xlarge",
		}
	case deployments.GCP:
		return []string{
			"n1-heighmem-2",
			"n1-heighmem-4",
			"n1-heighmem-8",
			"n1-heighmem-16",
			"n1-heighmem-32",
			"n1-heighmem-64",
			"n1-heighmem-96",
		}
	case deployments.Azure:
		return []string{}
	}

	pterm.Fatal.Printfln("'%s' does not support database sizes", platform)
	return []string{}
}
