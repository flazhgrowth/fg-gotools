package jwt

import (
	"errors"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken error = errors.New("invalid token")
)

type Claims struct {
	gojwt.RegisteredClaims
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewClaims(expiresAt time.Time, id string, username string, email string) Claims {
	return Claims{
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(expiresAt),
			IssuedAt:  gojwt.NewNumericDate(time.Now()),
		},
		ID:       id,
		Username: username,
		Email:    email,
	}
}
