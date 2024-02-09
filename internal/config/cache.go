package config

import (
	"time"
)
import "github.com/spf13/viper"

type CacheConfig struct {
	TTL               time.Duration `mapstructure:"ttl"`
	SearchExpiredTime time.Duration `mapstructure:"search_expired_interval"`
}

func InitCache() (*CacheConfig, error) {
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cacheConfig CacheConfig
	if err := viper.UnmarshalKey("cache", &cacheConfig); err != nil {
		return nil, err
	}
	return &cacheConfig, nil
}
