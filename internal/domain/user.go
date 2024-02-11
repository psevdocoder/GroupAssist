package domain

import "time"

type User struct {
	ID                int        `json:"id" db:"id"`
	Username          string     `json:"username" db:"username"`
	Name              string     `json:"name" db:"name"`
	PasswordHash      string     `json:"password_hash,omitempty" db:"password_hash"`
	Role              int        `json:"role" db:"role"`
	RegisteredAt      *time.Time `json:"registered_at,omitempty" db:"registered_at"`
	RegisterTokenHash string     `json:"register_token_hash,omitempty" db:"register_token"`
	IsActive          bool       `json:"is_active,omitempty" db:"is_active"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpInput struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required,min=12,max=48"`
	RegisterToken string `json:"register_token" binding:"required"`
}

type ApplyRegister struct {
	SignUpInput
	IsActive bool
}

type UpdateUserInput struct {
	Name string `json:"name"`
	Role int    `json:"role"`
}

type CreateUser struct {
	ID            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	Role          int    `json:"role" binding:"required,numeric,oneof=1 2 3"`
	RegisterToken string `json:"register_token"`
}
