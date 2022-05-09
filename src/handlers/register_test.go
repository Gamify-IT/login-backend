package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/register"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"testing"
	"time"
)

func TestRegisterUser_ShouldReturnOkIfUserNotAlreadyExists(t *testing.T) {
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

	hasher := &hash.Mock{FixedReturnValue: []byte("")}
	passwordHash, _ := hasher.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	stringPasswordHash := string(passwordHash)

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:           "user id",
			CreatedAt:    time.Now(),
			Name:         username,
			Email:        email,
			PasswordHash: stringPasswordHash,
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
	).Errors(db.ErrNotFound) // sets the object which should be returned in the function call

	mock.User.Expect(
		client.User.CreateOne(
			db.User.Name.Set(username),
			db.User.Email.Set(email),
			db.User.PasswordHash.Set(stringPasswordHash),
		),
	).Returns(expected)

	handler := RegisterUser(client, hasher)

	params := register.NewPostRegisterParams()
	params.Body = &models.Register{
		Password: &password,
		Username: &username,
		Email:    &email,
	}
	params.HTTPRequest = &http.Request{}

	result := handler(params)

	_, ok := result.(*register.PostRegisterOK)

	if !ok {
		t.Errorf("A valid registration should return a 200 status code")
	}
}

func TestRegisterUser_ShouldReturnBadRequestIfUserAlreadyExists(t *testing.T) {
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

	hasher := &hash.Bcrypt{}
	handler := RegisterUser(client, hasher)

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
