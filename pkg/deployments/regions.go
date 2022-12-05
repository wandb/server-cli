package deployments

func SupportedRegions(platform DeploymentPlatform) []string {
	switch platform {
	case AWS:
		return []string{
			"us-east-2",
			"us-east-1",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"eu-west-2",
			"us-gov-east-1",
			"us-gov-west-2",
		}
	case GCP:
		return []string{}
	}

	return []string{}
}
