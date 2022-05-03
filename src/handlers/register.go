package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/register"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

// registerUser let a user register with username, email and password
func registerUser(client *db.PrismaClient) register.PostRegisterHandlerFunc {
	return register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
		username := params.Body.Username
		email := params.Body.Email
		password := params.Body.Password

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

		existingUser, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil {
			return register.NewPostRegisterInternalServerError().WithPayload(&models.Error{
				Message: "Error finding existing user in database",
			})
		}

		if existingUser != nil {
			return register.NewPostRegisterBadRequest().WithPayload(&models.Error{
				Message: "Username already in use",
			})
		}

		client.User.CreateOne(db.User.Name.Set(*username), db.User.Email.Set(*email), db.User.PasswordHash.Set(string(hashedPassword)))

		createdUser, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil {
			return register.NewPostRegisterInternalServerError().WithPayload(&models.Error{
				Message: "Error finding registered user in database",
			})
		}

		return register.NewPostRegisterOK().WithPayload(&models.User{
			ID:   createdUser.ID,
			Name: createdUser.Name,
		})

	})
}
