package service

import "GroupAssist/internal/repository"

type Services struct {
	Subject *Subject
	Queue   *Queue
	User    *UserService
	Auth    *AuthService
}

func InitServices(repositories *repository.Repositories) *Services {
	return &Services{
		Subject: NewSubject(repositories.Subject),
		Queue:   NewQueue(repositories.Queue),
		User:    NewUserService(repositories.User),
		Auth:    NewAuthService(repositories.Auth),
	}
}
