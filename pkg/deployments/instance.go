package deployments

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"github.com/wandb/server-cli/pkg/api/deploy"
	"github.com/wandb/server-cli/pkg/config"
)

func GetInstanceContext() string {
	return viper.GetString("instance")
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

func (c *InstanceConfig) GetName() string {
	return c.name
}

func (c *InstanceConfig) GetLatestLicense() string {
	license, _ := deploy.GetLicense(c.GetDeploymentID())
	return license
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

func (c *InstanceConfig) SetInterface(path string, value interface{}) {
	viper.Set("instances."+c.name+"."+path, value)
}

func (c *InstanceConfig) GetInterface(path string, value interface{}) {
	viper.UnmarshalKey("instances."+c.name+"."+path, value)
}

func (c *InstanceConfig) Directory() string {
	path := filepath.Join(config.ConfigDir(), "instances", c.name)
	os.MkdirAll(path, os.ModePerm)
	return path
}

func (c *InstanceConfig) WriteFile(name string, context string) {
	path := filepath.Join(c.Directory(), name)

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	pterm.PrintOnError(err)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	pterm.Fatal.PrintOnError(err)

	defer f.Close()
	_, err = f.WriteString(context)
	pterm.Fatal.PrintOnError(err)
}

func (c *InstanceConfig) GetDeploymentID() string {
	return viper.GetString("instances." + c.name + ".deployment.id")
}

func (c *InstanceConfig) SetDeploymentID(value string) *InstanceConfig {
	viper.Set("instances."+c.name+".deployment.id", value)
	return c
}
