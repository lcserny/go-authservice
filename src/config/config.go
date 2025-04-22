package config

import (
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Host        string `mapstructure:"host"`
	ContextPath string `mapstructure:"contextPath"`
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/go-authservice/")
	viper.AddConfigPath("$HOME/.go-authservice")

	viper.AutomaticEnv()

	viper.SetDefault("port", 3003)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("contextPath", "/security")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logging.Warn("Config file not found; using defaults and environment variables")
		} else {
			panic("Error reading config file: " + err.Error())
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic("Unable to decode into struct: " + err.Error())
	}

	return &config
}
