package byob

type GoogleBYOBConfig struct {
	Zone      string
	ProjectID string
}

type AWSBYOBConfig struct{}

type BYOBConfig struct {
	BucketPrefix string
	Region       string
	Google       *GoogleBYOBConfig `yaml:",omitempty"`
	AWS          *AWSBYOBConfig    `yaml:",omitempty"`
}
