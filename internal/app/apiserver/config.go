package apiserver

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	BindAddr string
	LogLevel string
}

var config AppConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %s\n", err)
	}
}

func GetConfig() AppConfig {
	return config
}
