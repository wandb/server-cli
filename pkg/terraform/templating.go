package terraform

import (
	"context"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/pterm/pterm"
)

func getLastestVersionFromGithub() string {
	// githubClient := github.NewClient(nil)
	// res, _, err := githubClient.Repositories.GetLatestRelease(ctx, "hashicorp", "terraform")
	// if _, ok := err.(*github.RateLimitError); ok {
	// 	terraformVersionSpinner.Fail("Hit github rate limit.")
	// }
	// pterm.Fatal.PrintOnError(err)
	return ""
}

func getLatestVersionOfTerraform(ctx context.Context) string {
	terraformVersionSpinner, _ := pterm.DefaultSpinner.Start("Getting latest version of terraform")
	githubClient := github.NewClient(nil)
	res, _, err := githubClient.Repositories.GetLatestRelease(ctx, "hashicorp", "terraform")
	if _, ok := err.(*github.RateLimitError); ok {
		terraformVersionSpinner.Fail("Hit github rate limit.")
	}
	pterm.Fatal.PrintOnError(err)
	latestTerraformVersion := strings.TrimPrefix(*res.Name, "v")
	terraformVersionSpinner.Success("Latest version is v" + latestTerraformVersion)
	return latestTerraformVersion
}
