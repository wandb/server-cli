package terraform

import (
	_ "embed"

	"github.com/pterm/pterm"

	"github.com/wandb/server-cli/pkg/deployments"
	"github.com/wandb/server-cli/pkg/deployments/terraform/tfconfig"
)

func ConfigFlow(c tfconfig.TerraformConfig, platform string) {
	textInput := pterm.DefaultInteractiveTextInput
	confirmInput := pterm.DefaultInteractiveConfirm
	c.Namespace, _ = textInput.WithDefaultText("Namespace").Show()
	c.DomainName, _ = textInput.WithDefaultText("Domain name").Show()
	c.DeletionProtection = true
	c.Redis = true

	if c.Database == nil {
		c.Database = new(tfconfig.DatabaseConfig)
	}

	if platform == string(deployments.AWS) {
		if c.AWS == nil {
			c.AWS = new(tfconfig.AWSConfig)
		}
		c.DomainName, _ = textInput.WithDefaultText("Domain name").Show()
		hasACM, _ := confirmInput.
			WithDefaultValue(false).
			WithConfirmText("Do you have an existing ACM Certificate for TLS?").
			Show()
		if hasACM {
			c.AWS.ACMCertificateARN, _ = textInput.WithDefaultText("ACM Certifivate ARN").Show()
		}
	}

	if platform == string(deployments.Azure) {
		if c.Azure == nil {
			c.Azure = new(tfconfig.AzureConfig)
		}
	}
}

// func GenerateAWS(config *tfconfig.TerraformConfig) {
// 	b := new(strings.Builder)
// 	tmpl, _ := template.New("aws").Parse(privateCloudAWS)
// 	tmpl.Execute(b, config)
// 	fmt.Println(b.String())
// }

func NewTerraformConfig() *tfconfig.TerraformConfig {
	cfg := new(tfconfig.TerraformConfig)
	cfg.Database = new(tfconfig.DatabaseConfig)
	cfg.Docker = new(tfconfig.DockerConfig)
	cfg.Modules = new(tfconfig.Modules)
	return cfg
}
