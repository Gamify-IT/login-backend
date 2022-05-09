package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/register"
	"net/http"
	"testing"
	"time"
)

func TestRegisterUser_ShouldReturnBadRequestIfNoUserAlreadyExists(t *testing.T) {
	// create a new mock
	// this returns a mock prisma `client` and a `mock` object to set expectations
	client, mock, ensure := db.NewMock()
	// defer calling ensure, which makes sure all of the expectations were met and actually called
	// calling this makes sure that an error is returned if there was no query happening for a given expectation
	// and makes sure that all of them succeeded
	defer ensure(t)

	username := "Test User"
	email := "test@username.com"
	password := "password"
	passwordHash := "$2a$10$KIKrid5AyyXHKHXRt.zS7OrlYWqYv2FUJOXVOktCotczFKRhmVBW."

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           "user id",
			CreatedAt:    time.Now(),
			Name:         username,
			Email:        email,
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

	handler := RegisterUser(client)

	params := register.NewPostRegisterParams()
	params.Body = &models.Register{
		Password: &password,
		Username: &username,
		Email:    &email,
	}
	params.HTTPRequest = &http.Request{}

	result := handler(params)

	_, ok := result.(*register.PostRegisterBadRequest)

	if !ok {
		t.Errorf("A registration with already used username should return a 400 status code")
	}
}
