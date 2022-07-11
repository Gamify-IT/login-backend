package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/login"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"net/http"
	"testing"
	"time"
)

func TestLoginUser_ShouldReturnOKIfTheUserCredentialsAreValid(t *testing.T) {
	// Create a new database mock
	client, mock, ensure := db.NewMock()
	// At the end of the test, ensure that all of the expectations were met and actually called
	defer ensure(t)
	
	const testUserName = "Test User"

	// Prepare mock data
	password := "password"
	passwordHash := "$2a$10$KIKrid5AyyXHKHXRt.zS7OrlYWqYv2FUJOXVOktCotczFKRhmVBW."
	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           "test_user_id",
			CreatedAt:    time.Now(),
			Name:         testUserName,
			Email:        "test@username.com",
			PasswordHash: passwordHash,
		},
	}

	// Add expected query to mock database
	mock.User.Expect(
		client.User.FindFirst(
			db.User.Name.Equals(testUserName),
		),
	).Returns(expected)

	// Create test request parameters
	username := testUserName
	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &password,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	// Setup dependencies
	authenticator := auth.NewAuthenticator("secret", "token", time.Hour)
	hasher := &hash.Bcrypt{}

	// Test
	handler := LoginUser(client, authenticator, hasher)
	result := handler(params)

	// Check result
	okResult, ok := result.(*login.PostLoginOK)
	if !ok {
		t.Errorf("A valid login should return a 200 status code")
	}
	if okResult.Payload.Name != testUserName {
		t.Errorf("Expected username %q but got %q", testUserName, okResult.Payload.Name)
	}
	if okResult.Payload.ID != "test_user_id" {
		t.Errorf("Expected user ID %q but got %q", "test_user_id", okResult.Payload.ID)
	}
	auth.VerifyCookieHelper(t, authenticator, okResult.SetCookie)
}

func TestLoginUser_ShouldReturnBadRequestIfTheUserCredentialsAreNotValid(t *testing.T) {
	// Create a new database mock
	client, mock, ensure := db.NewMock()
	// At the end of the test, ensure that all of the expectations were met and actually called
	defer ensure(t)

	const testUserName = "Test User"

	// Prepare mock data
	wrongPassword := "wrong_password"
	passwordHash := "$2a$10$KIKrid5AyyXHKHXRt.zS7OrlYWqYv2FUJOXVOktCotczFKRhmVBW."
	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           "test_user_id",
			CreatedAt:    time.Now(),
			Name:         testUserName,
			Email:        "test@username.com",
			PasswordHash: passwordHash,
		},
	}

	// Add expected query to mock database
	mock.User.Expect(
		client.User.FindFirst(
			db.User.Name.Equals(testUserName),
		),
	).Returns(expected)

	// Setup dependencies
	authenticator := auth.NewAuthenticator("secret", "token", time.Hour)
	hasher := &hash.Bcrypt{}

	// Create test request parameters
	username := testUserName
	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &wrongPassword,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	// Test
	handler := LoginUser(client, authenticator, hasher)
	result := handler(params)

	// Check result
	if _, ok := result.(*login.PostLoginBadRequest); !ok {
		t.Errorf("A invalid login with wrong password should return a 400 status code")
	}
}

func TestLoginUser_ShouldReturnBadRequestIfTheUserDoesNotExists(t *testing.T) {
	// Create a new database mock
	client, mock, ensure := db.NewMock()
	// At the end of the test, ensure that all of the expectations were met and actually called
	defer ensure(t)

	const testUserName = "Test User"

	// Add expected query to mock database
	mock.User.Expect(
		client.User.FindFirst(
			db.User.Name.Equals(testUserName),
		),
	).Errors(db.ErrNotFound)

	// Setup dependencies
	authenticator := auth.NewAuthenticator("secret", "token", time.Hour)
	hasher := &hash.Bcrypt{}

	// Create test request parameters
	username := testUserName
	password := "password"
	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &password,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	// Test
	handler := LoginUser(client, authenticator, hasher)
	result := handler(params)

	// Check result
	if _, ok := result.(*login.PostLoginBadRequest); !ok {
		t.Errorf("A invalid login with non existing user should return a 400 status code")
	}
}
