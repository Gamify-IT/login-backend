package auth

import (
	"github.com/Gamify-IT/login-backend/src/helpers"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func TestAuthenticator_CookieName_ShouldReturnTheNameOfTheCookie(t *testing.T) {
	// Setup
	a := &Authenticator{cookieName: "cookie_name"}

	// Test
	r := a.CookieName()

	// Check result
	if r != "cookie_name" {
		t.Errorf("Expected cookie name %q but got %q", "cookie_name", r)
	}
}

func TestAuthenticator_Verify_ShouldFailOnAnEmptyString(t *testing.T) {
	// Setup
	a := &Authenticator{
		cookieName:       "cookie_name",
		secret:           "secret",
		validityDuration: time.Minute,
	}

	// Test
	cookie, err := a.Verify("")

	// Check result
	if err == nil {
		t.Errorf("Expected an error but got nil, cookie=%#v", cookie)
	}
}

func TestAuthenticator_Verify_ShouldFailOnAnExpiredToken(t *testing.T) {
	// Setup
	a := &Authenticator{
		cookieName:       "cookie_name",
		secret:           "secret",
		validityDuration: 0,
	}

	token, _ := a.GenerateJWT("test_user_id", "test_user_name")
	// let the token expire
	time.Sleep(time.Second)

	// Test
	result, err := a.Verify(token)

	// Check result
	if err == nil {
		t.Errorf("Expected an error but got nil, token=%#v, result=%#v", token, result)
	}
}

func TestAuthenticator_endToEnd(t *testing.T) {
	// Setup
	a := &Authenticator{
		cookieName:       "cookie_name",
		secret:           "secret",
		validityDuration: time.Minute,
	}

	// Test
	token, _ := a.GenerateJWT("test_user_id", "test_user_name")
	result, err := a.Verify(token)

	// Check result
	if err == nil {
		t.Errorf("Expected an error but got nil, token=%#v, result=%#v", token, result)
	}
	claims, ok := result.Claims.(jwt.MapClaims)
	if !ok {
		t.Errorf("Could not parse claims from result=%#v", result)
	}
	if id, ok := claims["id"].(string); ok {
		if id != "test_user_id" {
			t.Errorf("Expected id=%q but got %q", "test_user_id", id)
		}
	}
	if user, ok := claims["user"].(string); ok {
		if user != "test_user_id" {
			t.Errorf("Expected user=%q but got %q", "test_user_name", user)
		}
	}
}

func TestAuthenticator_GenerateTokenCookie_ShouldReturnAValidCookie(t *testing.T) {
	// Setup
	a := &Authenticator{
		cookieName:       "cookie_name",
		secret:           "secret",
		validityDuration: time.Minute,
	}

	// Test
	cookie, err := a.GenerateTokenCookie("test_user_id", "test_user_name")
	if err != nil {
		t.Error(err)
	}

	// Check result
	helpers.VerifyCookie(t, a, cookie)
}
