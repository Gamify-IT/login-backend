package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func NewAuthenticator(secret string) *Authenticator {
	return &Authenticator{secret: secret}
}

type Authenticator struct {
	secret string
}

func (j *Authenticator) GenerateJWT(id, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = name
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Authenticator) Verify(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Could not parse token: token method is %T instead of jwt.SigningMethodHMAC", token.Method)
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return err
	}

	return nil
}
