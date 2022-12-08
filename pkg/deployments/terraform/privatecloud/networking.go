package privatecloud

type NetworkingConfig struct {
	AllowedInboundCIDRs4 []string
	AllowedInboundCIDRs6 []string
	CreateNetwork        bool
}

func ConfigureNetwork() *NetworkingConfig {
	config := new(NetworkingConfig)
	config.CreateNetwork = true
	return new(NetworkingConfig)
}
