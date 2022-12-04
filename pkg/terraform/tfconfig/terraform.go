package tfconfig

type DockerConfig struct {
	Image   string `yaml:"image"`
	Version string `yaml:"version"`
}

type ModuleConfig struct {
	Version string `yaml:"version"`
}

type Modules struct {
	AWS        ModuleConfig `yaml:"aws"`
	Kubernetes ModuleConfig `yaml:"kubernetes"`
	Google     ModuleConfig `yaml:"google"`
	Azure      ModuleConfig `yaml:"azure"`
}

type DatabaseConfig struct {
	Version string `yaml:"version"`
	Size    string `yaml:"size"`
}

type AWSConfig struct {
	ACMCertificateARN string `yaml:"acm-certificate-arn"`
	ZoneID            string `yaml:"zone-id"`
}
type AzureConfig struct{}
type GoogleConfig struct{}

type TerraformConfig struct {
	Namespace          string          `yaml:"namespace"`
	Region             string          `yaml:"region"`
	License            string          `yaml:"license"`
	Redis              bool            `yaml:"redis"`
	Database           *DatabaseConfig `yaml:"database"`
	Docker             *DockerConfig   `yaml:"container"`
	Modules            *Modules        `yaml:"modules"`
	DeletionProtection bool            `yaml:"deletion-protection"`
	DomainName         string          `yaml:"domain-name"`

	AWS    *AWSConfig    `yaml:"aws"`
	Azure  *AzureConfig  `yaml:"azure"`
	Google *GoogleConfig `yaml:"google"`
}

func NewConfig() *TerraformConfig {
	cfg := new(TerraformConfig)
	cfg.Database = new(DatabaseConfig)
	cfg.Docker = new(DockerConfig)
	cfg.Modules = new(Modules)
	return cfg
}
