package terraform

import (
	"context"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/pterm/pterm"
)

func getLastestVersionFromGithub(owner string, repo string) string {
	ctx := context.Background()
	terraformVersionSpinner, _ := pterm.DefaultSpinner.Start("Getting version from Github")
	githubClient := github.NewClient(nil)
	res, _, err := githubClient.Repositories.GetLatestRelease(ctx, owner, repo)
	if _, ok := err.(*github.RateLimitError); ok {
		terraformVersionSpinner.Fail("Hit github rate limit.")
	}
	pterm.Fatal.PrintOnError(err)
	latestTerraformVersion := strings.TrimPrefix(*res.Name, "v")
	terraformVersionSpinner.Success("Latest version is v" + latestTerraformVersion)
	return ""
}

func GetLastestVersionOfTerraformAWS() string {
	return getLastestVersionFromGithub("wandb", "terraform-aws-wandb")
}

func GetLastestVersionOfTerraformGoogle() string {
	return getLastestVersionFromGithub("wandb", "terraform-google-wandb")
}

func GetLastestVersionOfTerraformAzure() string {
	return getLastestVersionFromGithub("wandb", "terraform-azure-wandb")
}
