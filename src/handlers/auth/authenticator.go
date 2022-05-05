package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateJWT(id, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = name
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString(os.Getenv("JWT_Key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
