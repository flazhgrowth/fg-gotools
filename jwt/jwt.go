package jwt

import (
	"encoding/json"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type JWT interface {
	GenerateToken(claims Claims, secretKey string) (string, error)
	ValidateToken(tokenString string, secretKey string) (claims Claims, err error)
}

type jwt struct{}

func NewJWT() JWT {
	return &jwt{}
}

func (j *jwt) GenerateToken(claims Claims, secretKey string) (string, error) {
	// Create a new token object
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwt) ValidateToken(tokenString string, secretKey string) (claims Claims, err error) {
	token, err := gojwt.Parse(tokenString, func(token *gojwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Error().Msgf("ValidateToken: %s", err.Error())
		return claims, ErrInvalidToken
	}

	if !token.Valid {
		log.Error().Msgf("ValidateToken.token.Valid: %v", token.Valid)
		return claims, ErrInvalidToken
	}

	mapClaims, ok := token.Claims.(gojwt.MapClaims)
	if !ok {
		log.Error().Msgf("ValidateToken.token.Claims.(Claims): %v", ok)
		return claims, ErrInvalidToken
	}
	byteMapClaims, err := json.Marshal(mapClaims)
	if err != nil {
		log.Error().Msgf("ValidateToken.json.Marshal(mapClaims): %s", err.Error())
		return claims, ErrInvalidToken
	}
	if err = json.Unmarshal(byteMapClaims, &claims); err != nil {
		log.Error().Msgf("ValidateToken.json.Unmarshal(byteMapClaims, &claims): %s", err.Error())
		return claims, ErrInvalidToken
	}

	return claims, nil
}
