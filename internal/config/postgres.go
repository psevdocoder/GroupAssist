package config

import (
	"github.com/spf13/viper"
	"log"
)

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"user"`
	DBName   string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl"`
	Password string `mapstructure:"password"`
}

func InitPostgres() *PostgresConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	var pgConfig PostgresConfig
	if err := viper.UnmarshalKey("postgres", &pgConfig); err != nil {
		log.Fatal(err)
	}

	return &pgConfig

}
