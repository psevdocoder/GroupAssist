package domain

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)
