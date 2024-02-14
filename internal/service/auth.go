package service

import (
	"GroupAssist/internal/config"
	"GroupAssist/internal/domain"
	"GroupAssist/pkg/bcrypt"
	"GroupAssist/pkg/randomStr"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthRepository interface {
	ApplyRegister(user domain.ApplyRegister) (domain.ResponseUser, error)
	GetByUsername(username string) (domain.User, error)
	GetByID(id int) (domain.User, error)
	GetRegisterToken(id int) (string, error)
	SetRefreshToken(userID int, refreshToken string, expiresAt time.Time, ip string) error
	GetRefreshToken(token string) (domain.Session, error)
}

type AuthService struct {
	repo       AuthRepository
	authConfig *config.Config
}

func NewAuthService(repo AuthRepository, conf *config.Config) *AuthService {
	return &AuthService{
		repo:       repo,
		authConfig: conf,
	}
}

func (a *AuthService) ApplyRegister(user domain.SignUpInput) (domain.ResponseUser, error) {
	hashedPassword, err := bcrypt.HashString(user.Password)
	if err != nil {
		return domain.ResponseUser{}, err
	}

	userInput := domain.ApplyRegister{
		SignUpInput: domain.SignUpInput{
			Username:      user.Username,
			Password:      hashedPassword,
			RegisterToken: user.RegisterToken,
		},
		IsActive: true,
	}
	return a.repo.ApplyRegister(userInput)
}

func (a *AuthService) CreateToken(input domain.SignInInput, ip string) (domain.SignInResponse, error) {
	user, err := a.repo.GetByUsername(input.Username)
	if err != nil {
		return domain.SignInResponse{}, domain.ErrInvalidCredentials
	}

	if ok := bcrypt.CheckStringHash(input.Password, user.PasswordHash); !ok {
		return domain.SignInResponse{}, domain.ErrInvalidCredentials
	}

	return a.generateJWT(user, ip)
}

func (a *AuthService) generateJWT(user domain.User, ip string) (domain.SignInResponse, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour + a.authConfig.Auth.AccessTokenTTL).Unix(),
	})

	accessToken, err := token.SignedString([]byte(a.authConfig.Jwt.SecretKey))
	if err != nil {
		return domain.SignInResponse{}, err
	}

	refreshToken := randomStr.GenerateRefreshToken()

	if err != nil {
		return domain.SignInResponse{}, err
	}

	expiresAt := time.Now().Add(time.Hour + a.authConfig.Auth.RefreshTokenTTL)

	if err = a.repo.SetRefreshToken(user.ID, refreshToken, expiresAt, ip); err != nil {
		return domain.SignInResponse{}, err
	}

	return domain.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthService) RefreshToken(token string) (domain.SignInResponse, error) {
	session, err := a.repo.GetRefreshToken(token)
	if err != nil {
		return domain.SignInResponse{}, err
	}

	if session.ExpiresAt.Before(time.Now()) {
		return domain.SignInResponse{}, domain.ErrRefreshTokenExpired
	}

	user, err := a.repo.GetByID(session.UserID)
	if err != nil {
		return domain.SignInResponse{}, err
	}

	return a.generateJWT(user, session.IP)
}
