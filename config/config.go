package config

import (
	"github.com/spf13/viper"
)
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}

func Error() error {
	configError := Error()
	return configError
}
