package authapi

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey = []byte("secert-string")

type Claims struct {
	Username string `json:string`
	jwt.StandardClaims
}

func verifyJwt(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("claims %s", claims)
	} else {
		return err
	}

	return nil
}

func newJwt(username string) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour)

	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func refreshJwt(tokenString string) (string, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	newToken, err := newJwt(claims.Username)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
