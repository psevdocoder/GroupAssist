package service

import "GroupAssist/internal/repository"

type Services struct {
	Subject *Subject
	Queue   *Queue
}

func InitServices(repositories *repository.Repositories) *Services {
	return &Services{
		Subject: NewSubject(repositories.Subject),
		Queue:   NewQueue(repositories.Queue),
	}
}
