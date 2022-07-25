// Package auth implements the logic required for user authentication using cookies.
package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

// NewAuthenticator returns a new Authenticator struct with all the required fields filled in.
//
// The secret is a symmetric key, used to generate and verify [JSON Web Tokens]
// The cookieName is the name of the authentication cookie.
// The validityDuration specifies the time after which the [JSON Web Tokens] expire and the cookie is deleted.
//
// [JSON Web Tokens]: https://jwt.io/
func NewAuthenticator(
	secret string,
	cookieName string,
	validityDuration time.Duration,
) *Authenticator {
	return &Authenticator{
		secret:           secret,
		cookieName:       cookieName,
		validityDuration: validityDuration,
	}
}

// Authenticator represents a security configuration including a cookieName, secret and validityDuration.
//
// You should create structs of this type using the [NewAuthenticator] function.
type Authenticator struct {
	cookieName       string
	secret           string
	validityDuration time.Duration
}

// CookieName returns the name of the authentication cookie.
func (a *Authenticator) CookieName() string {
	return a.cookieName
}

// GenerateTokenCookie returns a cookie header with a token containing the user id and corresponding display name.
//
// The cookie is valid for the configured validityDuration provided to [NewAuthenticator].
// It contains provided the user id and display name as payload.
func (a *Authenticator) GenerateTokenCookie(id, name string) (string, error) {
	token, err := a.GenerateJWT(id, name)

	if err != nil {
		return "", err
	}

	maxAge := int(a.validityDuration.Seconds())

	cookie := fmt.Sprintf("%s=%s; Path=/; Max-Age=%d; Secure; HttpOnly", a.cookieName, token, maxAge)

	return cookie, nil
}

// GenerateJWT returns a new JSON Web Token for the user.
// The token is valid for the configured validityDuration.
// It contains provided the user id and display name as payload.
func (a *Authenticator) GenerateJWT(id, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = name
	claims["exp"] = time.Now().Add(a.validityDuration).Unix()

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Verify that the token is valid.
func (a *Authenticator) Verify(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Could not parse token: token method is %T instead of jwt.SigningMethodHMAC", token.Method)
		}
		return []byte(a.secret), nil
	})
}
