package domain

import "time"

type SignUpInput struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required,alphanum,min=12,max=48"`
	RegisterToken string `json:"register_token" binding:"required"`
}

type ApplyRegister struct {
	SignUpInput
	IsActive bool
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type Session struct {
	ID           int64     `db:"id"`
	UserID       int       `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresAt    time.Time `db:"expires_at"`
	IP           string    `db:"ip_address"`
}
