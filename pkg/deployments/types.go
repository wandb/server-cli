package deployments

type DeploymentType string

const (
	ManagedDedicatedCloud DeploymentType = "W&B Managed Dedicated Cloud"
	ManagedPrivateCloud   DeploymentType = "W&B Managed Private Cloud"
	PrivateCloud          DeploymentType = "Self-Managed Private Cloud"
	BareMetal             DeploymentType = "Self-Managed Bare Metal"
)

var (
	deploymentTypesMap = map[string]DeploymentType{
		"W&B Managed Dedicated Cloud": ManagedDedicatedCloud,
		"W&B Managed Private Cloud":   ManagedPrivateCloud,
		"Self-Managed Private Cloud":  PrivateCloud,
		"Self-Managed Bare Metal":     BareMetal,
	}
)

func ParseType(str string) (DeploymentType, bool) {
	c, ok := deploymentTypesMap[str]
	return c, ok
}

type DeploymentPlatform string

const (
	AWS        DeploymentPlatform = "Amazon Web Services"
	GCP        DeploymentPlatform = "Google Cloud"
	Azure      DeploymentPlatform = "Azure"
	Host       DeploymentPlatform = "Host"
	Kubernetes DeploymentPlatform = "Kubernetes"
)

var (
	deploymentPlatformsMap = map[string]DeploymentPlatform{
		"Kubernetes":          Kubernetes,
		"Host":                Host,
		"Azure":               Azure,
		"GCP":                 GCP,
		"AWS":                 AWS,
		"Amazon Web Services": AWS,
	}
)

func ParsePlatform(str string) (DeploymentPlatform, bool) {
	c, ok := deploymentPlatformsMap[str]
	return c, ok
}

type DeploymentEngine string

const (
	BYOB      DeploymentEngine = "BYOB"
	HelmChart DeploymentEngine = "Helm Chart"
	Terraform DeploymentEngine = "Terraform"
	Docker    DeploymentEngine = "Docker"
)

var (
	deploymentEnginesMap = map[string]DeploymentEngine{
		"HelmChart": HelmChart,
		"Terraform": Terraform,
		"Docker":    Docker,
	}
)

func ParseEngine(str string) (DeploymentEngine, bool) {
	c, ok := deploymentEnginesMap[str]
	return c, ok
}
