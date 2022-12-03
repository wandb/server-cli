package config

import "github.com/spf13/viper"

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

func SwitchInstance(name string) *InstanceConfig {
	viper.Set("instance", name)
	viper.WriteConfig()
	return GetInstance()
}

func GetInstance() *InstanceConfig {
	ic := new(InstanceConfig)
	ic.name = viper.GetString("instance")
	return ic
}

type InstanceConfig struct {
	name string
}

func (c *InstanceConfig) Write() {
	viper.WriteConfig()
}

func (c *InstanceConfig) GetType() string {
	return viper.GetString("instances." + c.name + ".type")
}

func (c *InstanceConfig) SetType(value string) {
	viper.Set("instances."+c.name+".type", value)
}

func (c *InstanceConfig) GetAPIKey() string {
	return viper.GetString("instances." + c.name + ".apikey")
}

func (c *InstanceConfig) SetAPIKey(value string) {
	viper.Set("instances."+c.name+".apikey", value)
}

func (c *InstanceConfig) GetEngine() string {
	return viper.GetString("instances." + c.name + ".engine")
}

func (c *InstanceConfig) SetEngine(value string) {
	viper.Set("instances."+c.name+".engine", value)
}

func (c *InstanceConfig) GetPlatform() string {
	return viper.GetString("instances." + c.name + ".platform")
}

func (c *InstanceConfig) SetPlatform(value string) {
	viper.Set("instances."+c.name+".platform", value)
}

func (c *InstanceConfig) GetFQDN() string {
	return viper.GetString("instances." + c.name + ".fqdn")
}

func (c *InstanceConfig) SetFQDN(value string) {
	viper.Set("instances."+c.name+".fqdn", value)
}
