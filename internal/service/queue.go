package service

import "GroupAssist/internal/domain"

type QueueRepository interface {
	GetAllBySubject(id int) ([]domain.QueuesResponse, error)
	GetByID(id int) (domain.QueueResponse, error)
	Create(queue domain.Queue) (domain.Queue, error)
	Delete(id int) error
	Update(id int, queue domain.UpdateQueueInput) error
}

type Queue struct {
	repo QueueRepository
}

func NewQueue(repo QueueRepository) *Queue {
	return &Queue{
		repo: repo,
	}
}

func (q Queue) GetAllBySubject(id int) ([]domain.QueuesResponse, error) {
	return q.repo.GetAllBySubject(id)
}

func (q Queue) GetByID(id int) (domain.QueueResponse, error) {
	return q.repo.GetByID(id)
}

func (q Queue) Create(queue domain.Queue) (domain.Queue, error) {
	return q.repo.Create(queue)
}

func (q Queue) Delete(id int) error {
	return q.repo.Delete(id)
}

func (q Queue) Update(id int, queue domain.UpdateQueueInput) error {
	return q.repo.Update(id, queue)
}
