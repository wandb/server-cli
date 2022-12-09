package deployments

func SupportedRegions(platform DeploymentPlatform) []string {
	switch platform {
	case AWS:
		return []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"eu-central-1",
			"eu-west-1",
			"eu-south-1",
			"ap-southeast-2",
			"us-gov-east-1",
			"us-gov-west-2",
		}
	case GCP:
		return []string{
			"us-west1",
			"us-east1",
			"us-east4",
			"us-west1",
			"us-central1",
			"us-west2",
			"northamerica-northeast2",
			"australia-southeast1",
			"europe-west2",
			"europe-west3",
		}
	case Azure:
		return []string{
			"East US",
			"East US 2",
			"Central US",
			"North Central US",
			"South Central US",
			"West US",
			"West US 2",
		}
	}

	return []string{}
}
