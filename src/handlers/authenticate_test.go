package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/authenticate"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"net/http"
	"testing"
	"time"
)

func TestAuthenticateUser_ShouldReturnValidNewCookieIfLoggedIn(t *testing.T) {
	// Create JWT generator
	tokenGenerator := auth.NewAuthenticator("test secret", "token", time.Minute)

	// Create test request parameters
	testCookie, _ := tokenGenerator.GenerateTokenCookie("test_user_id", "test_user_name")
	params := authenticate.NewPostAuthenticateParams()
	params.HTTPRequest = &http.Request{}
	params.HTTPRequest.Header = http.Header{}
	params.HTTPRequest.Header.Add("Cookie", testCookie)

	// Test
	handler := AuthenticateUser(tokenGenerator)
	result := handler.Handle(params)

	// Check result
	if resultValue, ok := result.(*authenticate.PostAuthenticateOK); ok {
		auth.VerifyCookieHelper(t, tokenGenerator, resultValue.SetCookie)
	} else {
		t.Errorf("expected return type *authenticate.PostAuthenticateOK but got result: %#v", result)
	}
}

func TestAuthenticateUser_ShouldReturnUnauthorizedIfCookieNotSet(t *testing.T) {
	// Create JWT generator
	tokenGenerator := auth.NewAuthenticator("test secret", "token", time.Minute)

	// Create test request parameters without a Cookie header
	params := authenticate.NewPostAuthenticateParams()
	params.HTTPRequest = &http.Request{}
	params.HTTPRequest.Header = http.Header{}

	// Test
	handler := AuthenticateUser(tokenGenerator)
	result := handler.Handle(params)

	// Check result
	if _, ok := result.(*authenticate.PostAuthenticateUnauthorized); !ok {
		t.Errorf("expected return type *authenticate.PostAuthenticateOK but got result: %#v", result)
	}
}

func TestAuthenticateUser_ShouldRejectExpiredToken(t *testing.T) {
	// Create JWT generator
	tokenGenerator := auth.NewAuthenticator("test secret", "token", 0)

	// Create a cookie and let it expire
	testCookie, _ := tokenGenerator.GenerateTokenCookie("test_user_id", "test_user_name")
	time.Sleep(time.Second)

	// Create test request parameters
	params := authenticate.NewPostAuthenticateParams()
	params.HTTPRequest = &http.Request{}
	params.HTTPRequest.Header = http.Header{}
	params.HTTPRequest.Header.Add("Cookie", testCookie)

	// Test
	handler := AuthenticateUser(tokenGenerator)
	result := handler.Handle(params)

	// Check result
	if _, ok := result.(*authenticate.PostAuthenticateUnauthorized); !ok {
		t.Errorf("expected return type *authenticate.PostAuthenticateOK but got result: %#v", result)
	}
}
