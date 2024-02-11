package database

import (
	"GroupAssist/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConnection(config *config.Config) (*sqlx.DB, error) {

	pgConf := config.Postgres
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		pgConf.Host, pgConf.Port, pgConf.User, pgConf.DBName, pgConf.SSL, pgConf.Password))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
