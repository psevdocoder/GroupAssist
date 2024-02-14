package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server struct {
		Port int
	}

	Postgres struct {
		Host     string
		Port     int
		User     string
		DBName   string
		SSL      string
		Password string
	}

	Jwt struct {
		SecretKey string `envconfig:"secret_key"`
	}

	Cache struct {
		TTL               time.Duration `mapstructure:"ttl"`
		SearchExpiredTime time.Duration `mapstructure:"search_expired_interval"`
	}

	Auth struct {
		AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
		RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
	}
}

func New() (*Config, error) {
	var config Config
	if err := envconfig.Process("pg", &config.Postgres); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &config.Jwt); err != nil {
		return nil, err
	}

	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil

}
