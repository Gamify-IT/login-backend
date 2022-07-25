package handlers

import (
	"errors"
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/register"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser let a user sign up with email, password and a unique username.
func RegisterUser(client *db.PrismaClient, hasher hash.Hasher) register.PostRegisterHandlerFunc {
	return func(params register.PostRegisterParams) middleware.Responder {
		username := params.Body.Username
		email := params.Body.Email
		password := params.Body.Password

		hashedPassword, err := hasher.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

		if err != nil {
			return register.NewPostRegisterInternalServerError().WithPayload(&models.Error{
				Message: "An error occurred while encrypting your password",
			})
		}

		existingUser, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil && !errors.Is(err, db.ErrNotFound) {
			return register.NewPostRegisterInternalServerError().WithPayload(&models.Error{
				Message: "An error occurred while adding the user to the database",
			})
		} else if existingUser != nil {
			return register.NewPostRegisterBadRequest().WithPayload(&models.Error{
				Message: "Username already in use",
			})
		}

		createdUser, err := client.User.CreateOne(
			db.User.Name.Set(*username),
			db.User.Email.Set(*email),
			db.User.PasswordHash.Set(string(hashedPassword)),
		).Exec(params.HTTPRequest.Context())

		if err != nil {
			return register.NewPostRegisterInternalServerError().WithPayload(&models.Error{
				Message: "An error occurred while adding the user to the database",
			})
		}

		return register.NewPostRegisterOK().WithPayload(&models.User{
			ID:   createdUser.ID,
			Name: createdUser.Name,
		})
	}
}
