package config

import "github.com/spf13/viper"

func GetInstanceContext() string {
	return viper.GetString("context")
}

func SetInstanceContext(key string, value interface{}) {
	viper.Set("instance."+GetInstanceContext()+"."+key, value)
}
