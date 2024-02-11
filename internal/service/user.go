package service

import (
	"GroupAssist/internal/domain"
	register_tokens "GroupAssist/pkg/invite-tokens"
)

type UserRepository interface {
	Create(user domain.CreateUser) (domain.CreateUser, error)
	Update(id int, user domain.UpdateUserInput) error
	Delete(id int) error
	GetAll() ([]domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(createUser domain.CreateUser) (domain.CreateUser, error) {
	createUser.RegisterToken = register_tokens.GenerateRegisterToken(32)
	return u.repo.Create(createUser)
}

func (u *UserService) Update(id int, user domain.UpdateUserInput) error {
	return u.repo.Update(id, user)
}

func (u *UserService) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UserService) GetAll() ([]domain.User, error) {
	return u.repo.GetAll()
}
