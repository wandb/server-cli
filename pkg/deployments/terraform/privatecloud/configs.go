package privatecloud

import (
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/deployments"
)

type DockerConfig struct {
	Version string
	Image   string
}

type PrivateCloudConfig struct {
	APIKey string `yaml:"-"`

	// Azure
	AzureSubscriptionID string `yaml:",omitempty"`

	// AWS
	ACMCertificateARN string `yaml:",omitempty"`

	// Google
	GoogleProjectID string `yaml:",omitempty"`

	// Terraform
	ModuleKubeVersion   string `yaml:"-"`
	ModuleGoogleVersion string `yaml:"-"`
	ModuleAzureVersion  string `yaml:"-"`
	ModuleHelmVersion   string `yaml:"-"`

	// General
	Region             string
	Namespace          string
	EnableRedis        bool
	DeletionProtection bool
	UseInternalQueue   bool

	// Instance Properties
	DeploymentID string `yaml:"-"`
	License      string `yaml:"-"`

	// DNS
	Subdomain    string
	DomainName   string
	ExternalDNS  bool
	PublicAccess bool

	// Loadbalancer
	LoadBalancer *LoadBalancerConfig

	// Networking

	// Kube-cluster
	KubernetesPublicAccess bool
	KubernetesVMSize       bool

	// Docker Config
	Docker *DockerConfig `yaml:"-"`

	OIDCIssuer        string
	OIDCClientID      string
	OIDCSecret        string
	OIDCAuthMethod    string
	DisableCodeSaving bool

	// Database
	Database *DatabaseConfig
}

func ConfigurePrivateCloud() {
	i := deployments.GetInstance()
	platform := i.GetPlatform()

	config := new(PrivateCloudConfig)

	config.Docker = new(DockerConfig)
	config.Docker.Version = ""
	config.Docker.Image = "wandb/local"

	config.APIKey = viper.GetString("wandb.apikey")
	config.DeploymentID = i.GetDeploymentID()
	config.License = i.GetLatestLicense()

	config.DeletionProtection = true
	config.EnableRedis = true
	config.UseInternalQueue = true
	config.DisableCodeSaving = false

	config.Database = ConfigureDatabase(platform)
	config.LoadBalancer = ConfigureLoadBalancer()

	i.SetInterface("private-cloud", config)
	i.Write()
}
