package deployments

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/config"
	"github.com/wandb/server-cli/pkg/deployments/terraform/tfconfig"
)

func GetInstanceContext() string {
	return viper.GetString("context")
}

func SetInstanceContext(key string, value interface{}) {
	viper.Set("instance."+GetInstanceContext()+"."+key, value)
}

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

func (c *InstanceConfig) GetType() DeploymentType {
	t, _ := ParseType(viper.GetString("instances." + c.name + ".type"))
	return t
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

func (c *InstanceConfig) GetEngine() DeploymentEngine {
	t, _ := ParseEngine(viper.GetString("instances." + c.name + ".engine"))
	return t
}

func (c *InstanceConfig) SetEngine(value string) {
	viper.Set("instances."+c.name+".engine", value)
}

func (c *InstanceConfig) GetPlatform() DeploymentPlatform {
	p, _ := ParsePlatform(viper.GetString("instances." + c.name + ".platform"))
	return p
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

func (c *InstanceConfig) SetTerraformConfig(value *tfconfig.TerraformConfig) {
	viper.Set("instances."+c.name+".terraform", value)
}

func (c *InstanceConfig) GetTerraformConfig(value tfconfig.TerraformConfig) tfconfig.TerraformConfig {
	config := tfconfig.NewConfig()
	viper.UnmarshalKey("instances."+c.name+".terraform", config)
	return *config
}

func (c *InstanceConfig) InstanceDirectory() string {
	path := filepath.Join(config.ConfigDir(), "instances")
	os.MkdirAll(path, os.ModePerm)
	return path
}

func (c *InstanceConfig) GetDeploymentID() string {
	return viper.GetString("instances." + c.name + ".deployment.id")
}

func (c *InstanceConfig) SetDeploymentID(value string) *InstanceConfig {
	viper.Set("instances."+c.name+".deployment.id", value)
	return c
}
