package repository

import (
	"GroupAssist/internal/repository/psql"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Subject *psql.SubjectRepository
	Queue   *psql.QueueRepository
	User    *psql.UserRepository
	Auth    *psql.AuthRepository
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Subject: psql.NewSubjectRepository(db),
		Queue:   psql.NewQueueRepository(db),
		User:    psql.NewUserRepository(db),
		Auth:    psql.NewAuthRepository(db),
	}
}
