package privatecloud

import (
	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/deployments"
)

type GeneralConfig struct {
	Namespace        string
	Region           string
	EnableRedis      bool
	UseInternalQueue bool
}

func GetCloudRegion(platform deployments.DeploymentPlatform) []string {
	switch platform {
	case deployments.AWS:
		return []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"eu-central-1",
			"eu-west-1",
			"eu-west-1",
			"ap-southeast-1",
			"ap-southeast-2",
			"ap-northeast-1",
		}
	case deployments.GCP:
		return []string{}
	case deployments.Azure:
		return []string{}
	}

	pterm.Fatal.Printfln("'%s' Still not available", platform)
	return []string{}
}

func NamespaceConfig(config *GeneralConfig) {
	petName := deployments.PetName()

	pterm.DefaultParagraph.
		Println("The namespace will prefix most of W&B resources name that will be created.\n" +
			"If not set a random name will be generated and used")

	namespace, _ := pterm.DefaultInteractiveTextInput.Show("Namespace")
	if len(namespace) != 0 {
		config.Namespace = namespace
		pterm.Println(namespace)
		return
	}
	pterm.Info.Printf("No namespace was informed, the default '%s' will be used\n", petName)

	config.Namespace = petName
}

func RegionConfig(config *GeneralConfig, platform deployments.DeploymentPlatform) {
	regions := GetCloudRegion(platform)

	region, _ := pterm.DefaultInteractiveSelect.WithOptions(regions).Show("Select AWS Region ")
	config.Region = region
}

func EnableRedisConfig(config *GeneralConfig) {
	enable, _ := pterm.DefaultInteractiveConfirm.Show("Would you like to enable external Redis?")
	if !enable {
		return
	}

	config.EnableRedis = enable
}

func UseInternalQueue(config *GeneralConfig) {
	enable, _ := pterm.DefaultInteractiveConfirm.Show("Would you like to enable external queue (SNS)?")
	if !enable {
		return
	}

	config.EnableRedis = enable
}

func GeneralConfiguration(platform deployments.DeploymentPlatform) *GeneralConfig {
	config := new(GeneralConfig)
	NamespaceConfig(config)
	RegionConfig(config, platform)
	EnableRedisConfig(config)
	UseInternalQueue(config)
	return config
}
