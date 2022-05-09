package handlers

import (
	"errors"
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/login"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

// loginUser let a user log in
func LoginUser(client *db.PrismaClient, generator *auth.Authenticator) login.PostLoginHandlerFunc {
	return login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
		username := params.Body.Username
		password := params.Body.Password

		user, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil && !errors.Is(err, db.ErrNotFound) {
			return login.NewPostLoginInternalServerError().WithPayload(&models.Error{
				Message: "Error finding existing user in database",
			})
		}
		if errors.Is(err, db.ErrNotFound) {
			return login.NewPostLoginBadRequest()
		}
		passwordCompareError := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*password))
		if passwordCompareError != nil {
			return login.NewPostLoginBadRequest()
		}

		token, err := generator.GenerateJWT(user.ID, user.Name)

		if err != nil {
			return login.NewPostLoginInternalServerError().WithPayload(&models.Error{
				Message: "Error creating token",
			})
		}

		return login.NewPostLoginOK().WithPayload(&models.LoginSuccess{
			ID:    user.ID,
			Name:  user.Name,
			Token: token,
		})
	})
}
