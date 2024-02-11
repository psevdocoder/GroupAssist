package service

import (
	"GroupAssist/internal/domain"
	"GroupAssist/pkg/hashing"
	"log"
)

type AuthRepository interface {
	ApplyRegister(user domain.ApplyRegister) (domain.User, error)
	GetByCredentials(input domain.SignInInput) (domain.User, error)
	GetRegisterTokenHash(id int) (string, error)
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) ApplyRegister(user domain.SignUpInput) (domain.User, error) {
	hashedPassword, err := hashing.HashString(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	userInput := domain.ApplyRegister{
		SignUpInput: domain.SignUpInput{
			Username:      user.Username,
			Password:      hashedPassword,
			RegisterToken: user.RegisterToken,
		},
		IsActive: true,
	}
	log.Printf("userInput: %+v", userInput)
	return a.repo.ApplyRegister(userInput)
}

func (a *AuthService) GetByCredentials(input domain.SignInInput) (domain.User, error) {
	return a.repo.GetByCredentials(input)
}
