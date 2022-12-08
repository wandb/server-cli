package privatecloud

type DatabaseConfig struct {
	Version string
	Size    string
}

func ConfigureDatabase() *DatabaseConfig {
	return new(DatabaseConfig)
}
