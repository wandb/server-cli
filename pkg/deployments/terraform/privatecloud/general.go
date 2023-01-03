package privatecloud

import (
	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/deployments"
)

type GeneralConfig struct {
	Namespace          string
	Region             string
	EnableRedis        bool
	UseInternalQueue   bool
	DeletionProtection bool
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
	var petName string

	pterm.DefaultParagraph.
		Println("The namespace will prefix most of W&B resources name that will be created.\n" +
			"If not set a random name will be generated and used")

	namespace, _ := pterm.DefaultInteractiveTextInput.Show("Namespace")
	if len(namespace) != 0 {
		config.Namespace = namespace
		pterm.Println(namespace)
		return
	}
	petName = deployments.PetName()
	pterm.Info.Printf("No namespace was informed, the default '%s' will be used\n", petName)

	config.Namespace = petName
}

func RegionConfig(config *GeneralConfig, platform deployments.DeploymentPlatform) {
	regions := GetCloudRegion(platform)

	region, _ := pterm.DefaultInteractiveSelect.WithOptions(regions).Show("Select Cloud Region")
	config.Region = region
}

func EnableRedisConfig(config *GeneralConfig) {
	enable, _ := pterm.DefaultInteractiveConfirm.Show("Would you like to enable external Redis?")

	config.EnableRedis = enable
}

func UseInternalQueueConfig(config *GeneralConfig) {
	enable, _ := pterm.DefaultInteractiveConfirm.Show("Would you like to enable external queue (SNS)?")

	config.EnableRedis = enable
}

func DeletionProtectionConfig(config *GeneralConfig) {
	enable, _ := pterm.DefaultInteractiveConfirm.
		Show("Would you like to DISABLE Object Storage Deletion Protection?")
	if enable {
		doubleCheck, _ := pterm.DefaultInteractiveConfirm.
			WithTextStyle(pterm.Warning.MessageStyle).
			Show("You're about to disable Object Storage Deletion Protection, are you sure?")
		pterm.Println()
		if doubleCheck {
			pterm.Warning.Println("Object Storage Deletion Protection DISABLED.")
			pterm.Println()
			config.DeletionProtection = doubleCheck
		} else {
			config.DeletionProtection = enable
		}
	}
}

func GeneralConfiguration(platform deployments.DeploymentPlatform) *GeneralConfig {
	config := new(GeneralConfig)
	NamespaceConfig(config)
	RegionConfig(config, platform)
	EnableRedisConfig(config)
	UseInternalQueueConfig(config)
	DeletionProtectionConfig(config)
	return config
}
