package domain

import "errors"

var (
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrRefreshTokenExpired    = errors.New("refresh token expired")
	ErrInvalidToken           = errors.New("invalid token")
	ErrInvalidSignature       = errors.New("invalid token signature")
	ErrUnhandledToken         = errors.New("couldn't handle this token")
	ErrUnexpectedSigningToken = errors.New("unexpected signing method")
)
