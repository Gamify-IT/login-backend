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
	// create a new mock
	// this returns a mock prisma `client` and a `mock` object to set expectations
	client, mock, ensure := db.NewMock()
	// defer calling ensure, which makes sure all of the expectations were met and actually called
	// calling this makes sure that an error is returned if there was no query happening for a given expectation
	// and makes sure that all of them succeeded
	defer ensure(t)

	id := "user id"
	username := "Test User"
	password := "password"
	passwordHash := "$2a$10$KIKrid5AyyXHKHXRt.zS7OrlYWqYv2FUJOXVOktCotczFKRhmVBW."

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           id,
			CreatedAt:    time.Now(),
			Name:         username,
			Email:        "test@username.com",
			PasswordHash: passwordHash,
		},
	}

	// start the expectation
	mock.User.Expect(
		// define your exact query as in your tested function
		// call it with the exact arguments which you expect the function to be called with
		// you can copy and paste this from your tested function, and just put specific values into the arguments
		client.User.FindFirst(
			db.User.Name.Equals(username),
		),
	).Returns(expected) // sets the object which should be returned in the function call

	authenticator := auth.NewAuthenticator("asdf", time.Hour)
	hasher := &hash.Bcrypt{}

	handler := LoginUser(client, authenticator, hasher)

	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &password,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	result := handler(params)

	okResult, ok := result.(*login.PostLoginOK)

	if !ok {
		t.Errorf("A valid login should return a 200 status code")
	}

	if okResult.Payload.Name != username {
		t.Errorf("Expected username %q but got %q", username, okResult.Payload.Name)
	}

	if okResult.Payload.ID != id {
		t.Errorf("Expected user ID %q but got %q", id, okResult.Payload.ID)
	}
}

func TestLoginUser_ShouldReturnBadRequestIfTheUserCredentialsAreNotValid(t *testing.T) {
	// create a new mock
	// this returns a mock prisma `client` and a `mock` object to set expectations
	client, mock, ensure := db.NewMock()
	// defer calling ensure, which makes sure all of the expectations were met and actually called
	// calling this makes sure that an error is returned if there was no query happening for a given expectation
	// and makes sure that all of them succeeded
	defer ensure(t)

	username := "Test User"
	password := "passwordThatNotMatches"
	passwordHash := "$2a$10$KIKrid5AyyXHKHXRt.zS7OrlYWqYv2FUJOXVOktCotczFKRhmVBW."

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           "user id",
			CreatedAt:    time.Now(),
			Name:         username,
			Email:        "test@username.com",
			PasswordHash: passwordHash,
		},
	}

	// start the expectation
	mock.User.Expect(
		// define your exact query as in your tested function
		// call it with the exact arguments which you expect the function to be called with
		// you can copy and paste this from your tested function, and just put specific values into the arguments
		client.User.FindFirst(
			db.User.Name.Equals(username),
		),
	).Returns(expected) // sets the object which should be returned in the function call

	authenticator := auth.NewAuthenticator("asdf", time.Hour)
	hasher := &hash.Bcrypt{}

	handler := LoginUser(client, authenticator, hasher)

	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &password,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	result := handler(params)

	_, ok := result.(*login.PostLoginBadRequest)

	if !ok {
		t.Errorf("A invalid login with wrong password should return a 400 status code")
	}
}

func TestLoginUser_ShouldReturnBadRequestIfTheUserDoesNotExists(t *testing.T) {
	// create a new mock
	// this returns a mock prisma `client` and a `mock` object to set expectations
	client, mock, ensure := db.NewMock()
	// defer calling ensure, which makes sure all of the expectations were met and actually called
	// calling this makes sure that an error is returned if there was no query happening for a given expectation
	// and makes sure that all of them succeeded
	defer ensure(t)

	username := "Test User"
	password := "password"

	// start the expectation
	mock.User.Expect(
		// define your exact query as in your tested function
		// call it with the exact arguments which you expect the function to be called with
		// you can copy and paste this from your tested function, and just put specific values into the arguments
		client.User.FindFirst(
			db.User.Name.Equals(username),
		),
	).Errors(db.ErrNotFound) // sets the object which should be returned in the function call

	authenticator := auth.NewAuthenticator("asdf", time.Hour)
	hasher := &hash.Bcrypt{}

	handler := LoginUser(client, authenticator, hasher)

	params := login.NewPostLoginParams()
	params.Body = &models.Login{
		Password: &password,
		Username: &username,
	}
	params.HTTPRequest = &http.Request{}

	result := handler(params)

	_, ok := result.(*login.PostLoginBadRequest)

	if !ok {
		t.Errorf("A invalid login with non existing user should return a 400 status code")
	}
}
