package util

import (
	"github.com/spf13/viper"
)

// Config stores all configurations of the application
type Config struct {
	SERVER_ADDRESS      string `mapstructure:"SERVER_ADDRESS"`
	AUTH_SYSTEM_ADDRESS string `mapstructure:"AUTH_SYSTEM_ADDRESS"`
	MAPS_SYSTEM_ADDRESS string `mapstructure:"MAPS_SYSTEM_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
