package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	SSL      string
	Password string
}

func InitPostgres() (*PostgresConfig, error) {
	var pgConfig PostgresConfig
	if err := envconfig.Process("pg", &pgConfig); err != nil {
		return nil, err
	}
	log.Printf("Postgres config: %+v\n", pgConfig)
	return &pgConfig, nil
}
