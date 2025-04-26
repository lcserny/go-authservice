package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/spf13/viper"
)

type AuthenticationConfig struct {
	Algorithm               string   `mapstructure:"algorithm" validate:"oneof=HS256 HS384 HS512 RS256 RS384 RS512 ES256 ES384 ES512 PS256 PS384 PS512 none"`
	PrivateKey              string   `mapstructure:"privateKey"`
	PublicKey               string   `mapstructure:"publicKey"`
	Secret                  string   `mapstructure:"secret"`
	AccessExpirationMinutes int      `mapstructure:"accessExpirationMinutes"`
	RefreshExpirationDays   int      `mapstructure:"refreshExpirationDays"`
	RefreshTokenName        string   `mapstructure:"refreshTokenName"`
	Issuer                  string   `mapstructure:"issuer"`
	Audience                []string `mapstructure:"audience"`
	Salt                    string   `mapstructure:"salt"`
}

type DatabaseConfig struct {
	Url        string `mapstructure:"url"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	AuthSource string `mapstructure:"authSource"`
	Type       string `mapstructure:"type" validate:"oneof=postgres mariadb sqlite mongodb"`
}

type LogConfig struct {
	File    string `mapstructure:"file"`
	Level   string `mapstructure:"level"`
	Json    bool   `mapstructure:"json"`
	Size    string `mapstructure:"size"`
	NrFiles string `mapstructure:"nrFiles"`
}

type ApplicationConfig struct {
	Name string    `mapstructure:"name"`
	Port int       `mapstructure:"port"`
	Path string    `mapstructure:"path"`
	Env  string    `mapstructure:"env"`
	Log  LogConfig `mapstructure:"log"`
}

type Config struct {
	Application    ApplicationConfig    `mapstructure:"application"`
	Database       DatabaseConfig       `mapstructure:"database"`
	Authentication AuthenticationConfig `mapstructure:"authentication"`
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/go-authservice/")
	viper.AddConfigPath("$HOME/.go-authservice")

	viper.AutomaticEnv()

	viper.SetDefault("application.log.size", "20m")
	viper.SetDefault("application.log.nrFiles", "14d")
	viper.SetDefault("database.authSource", "admin")
	viper.SetDefault("database.type", "mongodb")
	viper.SetDefault("authentication.saltTimes", 10)
	viper.SetDefault("authentication.refreshTokenName", "refreshToken")

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

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		panic("Configuration validation failed: " + err.Error())
	}

	return &config
}
