package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	//viper.SetConfigFile("D:\\goProject\\reading-notes\\config\\config.properties")
	viper.AddConfigPath("./config")
	error := viper.ReadInConfig()
	if error != nil {
		logrus.Error(error)
	}
}

func Error() error {
	configError := Error()
	return configError
}
