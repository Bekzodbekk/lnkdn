package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateTokens(id string, secretKey string) (string, error) {
	accessTokenClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
