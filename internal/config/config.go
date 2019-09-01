package config

import (
	"log"

	"github.com/spf13/viper"
)

type Constants struct {
	SplashHost string
	SplashPort string
}

type Config struct {
	Constants
}


func InitConfig() *Config {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants

	if err != nil {
		log.Panic("There was an error setting the configuration")
	}

	return &config
}

func initViper () (Constants, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/go-splash")
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
	viper.SetDefault("SplashHost", "localhost")
	viper.SetDefault("SplashPort", "8050")

	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
