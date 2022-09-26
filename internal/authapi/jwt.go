package authapi

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func newJwt(username string) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour)

	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	tokenString, err := token.SignedString([]byte("secret-string"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
