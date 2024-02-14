package domain

import "time"

type User struct {
	ID            int        `json:"id" db:"id"`
	Username      string     `json:"username" db:"username"`
	Name          string     `json:"name" db:"name"`
	PasswordHash  string     `json:"password_hash,omitempty" db:"password_hash"`
	Role          int        `json:"role" db:"role"`
	RegisteredAt  *time.Time `json:"registered_at,omitempty" db:"registered_at"`
	RegisterToken string     `json:"register_token_hash,omitempty" db:"register_token"`
	IsActive      bool       `json:"is_active,omitempty" db:"is_active"`
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

type ResponseUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}
