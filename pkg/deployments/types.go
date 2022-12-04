package deployments

type DeploymentType string

const (
	ManagedDedicatedCloud DeploymentType = "W&B Managed Dedicated Cloud"
	ManagedPrivateCloud   DeploymentType = "W&B Managed Private Cloud"
	PrivateCloud          DeploymentType = "Self-Managed Private Cloud"
	BareMetal             DeploymentType = "Self-Managed Bare Metal"
)

type DeploymentPlatform string

const (
	AWS        DeploymentPlatform = "Amazon Web Services"
	GCP        DeploymentPlatform = "Google Cloud"
	Azure      DeploymentPlatform = "Azure"
	Host       DeploymentPlatform = "Host"
	Kubernetes DeploymentPlatform = "Kubernetes"
)

type DeploymentEngine string

const (
	BYOB      DeploymentEngine = "BYOB"
	HelmChart DeploymentEngine = "Helm Chart"
	Terraform DeploymentEngine = "Terraform"
	Docker    DeploymentEngine = "Docker"
)
