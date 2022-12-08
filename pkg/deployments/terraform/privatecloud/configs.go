package privatecloud

import (
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/deployments"
)

type PrivateCloudConfig struct {
	APIKey string `yaml:",omit"`

	// Azure
	AzureSubscriptionID string `yaml:",omitempty"`

	// AWS
	ACMCertificateARN string `yaml:",omitempty"`

	// Google
	GoogleProjectID string `yaml:",omitempty"`

	// Terraform
	ModuleKubeVersion   string `yaml:",omit"`
	ModuleGoogleVersion string `yaml:",omit"`
	ModuleAzureVersion  string `yaml:",omit"`
	ModuleHelmVersion   string `yaml:",omit"`

	// General
	Region             string
	Namespace          string
	EnableRedis        bool
	DeletionProtection bool
	UseInternalQueue   bool

	// Instance Properties
	DeploymentID string `yaml:",omit"`
	License      string `yaml:",omit"`

	// DNS
	Subdomain    string
	DomainName   string
	ExternalDNS  bool
	PublicAccess bool

	// Loadbalancer
	AllowedInboundCIDRs4 []string
	AllowedInboundCIDRs6 []string

	// Networking

	// Kube-cluster
	KubernetesPublicAccess bool
	KubernetesVMSize       bool

	// Docker Config
	DockerImage   string
	DockerVersion string

	OIDCIssuer        string
	OIDCClientID      string
	OIDCSecret        string
	OIDCAuthMethod    string
	DisableCodeSaving bool

	// Database
	DatabaseVersion string
	DatabaseSize    string
}

func ConfigurePrivateCloud() {
	i := deployments.GetInstance()
	config := new(PrivateCloudConfig)

	config.APIKey = viper.GetString("wandb.apikey")
	config.DeploymentID = i.GetDeploymentID()
	config.License = i.GetLatestLicense()
	config.DeletionProtection = true
	config.EnableRedis = true
	config.UseInternalQueue = true
	config.DisableCodeSaving = false
}
