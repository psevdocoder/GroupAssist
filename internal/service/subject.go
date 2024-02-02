package service

import "GroupAssist/internal/domain"

type SubjectRepository interface {
	GetAll() ([]domain.Subject, error)
	GetByID(id int) (domain.Subject, error)
	Create(subject domain.Subject) (domain.Subject, error)
	Delete(id int) error
	Update(id int, subject domain.UpdateSubjectInput) error
}

type Subject struct {
	repo SubjectRepository
}

func NewSubject(repo SubjectRepository) *Subject {
	return &Subject{
		repo: repo,
	}
}

func (s *Subject) GetAll() ([]domain.Subject, error) {
	return s.repo.GetAll()
}

func (s *Subject) GetByID(id int) (domain.Subject, error) {
	return s.repo.GetByID(id)
}

func (s *Subject) Create(subject domain.Subject) (domain.Subject, error) {
	return s.repo.Create(subject)
}

func (s *Subject) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Subject) Update(id int, subject domain.UpdateSubjectInput) error {
	return s.repo.Update(id, subject)
}
