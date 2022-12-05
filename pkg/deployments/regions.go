package deployments

func SupportedRegions(platform string) []string {
	switch platform {
	case string(AWS):
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
	case string(GCP):
		return []string{}
	}

	return []string{}
}
