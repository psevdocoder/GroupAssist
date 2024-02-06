package config

import (
	"github.com/joho/godotenv"
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

func InitPostgres() *PostgresConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var pgConfig PostgresConfig
	if err := envconfig.Process("pg", &pgConfig); err != nil {
		log.Fatal(err)
	}
	log.Printf("Postgres config: %+v\n", pgConfig)

	return &pgConfig

}
