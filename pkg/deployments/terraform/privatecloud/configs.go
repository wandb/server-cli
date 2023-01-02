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
	General            *GeneralConfig
	Region             string
	Namespace          string
	EnableRedis        bool
	UseInternalQueue   bool
	DeletionProtection bool

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

	// General configuration (Should I creatre General Section instead?)
	config.General = GeneralConfiguration(platform)
	config.Namespace = config.General.Namespace
	config.Region = config.General.Region
	config.EnableRedis = config.General.EnableRedis
	config.UseInternalQueue = config.General.UseInternalQueue
	config.DeletionProtection = true

	config.DisableCodeSaving = false
	config.Database = ConfigureDatabase(platform)
	config.LoadBalancer = ConfigureLoadBalancer()

	i.SetInterface("private-cloud", config)
	i.Write()
}
