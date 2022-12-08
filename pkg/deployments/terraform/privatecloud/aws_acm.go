package privatecloud

type AWSACMConfig struct {
	ACMCertificateARN string
}

func ConfigureAWSACM() *AWSACMConfig {
	return new(AWSACMConfig)
}
