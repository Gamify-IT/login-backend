package handlers

import (
	"errors"
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/login"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

// LoginUser let a user log in
func LoginUser(client *db.PrismaClient, generator *auth.Authenticator, hash hash.Hasher) login.PostLoginHandlerFunc {
	return func(params login.PostLoginParams) middleware.Responder {
		username := params.Body.Username
		password := params.Body.Password

		user, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil && !errors.Is(err, db.ErrNotFound) {
			log.Println(err)
			return login.NewPostLoginInternalServerError().WithPayload(&models.Error{
				Message: "Error finding existing user in database",
			})
		}
		if errors.Is(err, db.ErrNotFound) {
			return login.NewPostLoginBadRequest()
		}
		passwordCompareError := hash.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*password))
		if passwordCompareError != nil {
			return login.NewPostLoginBadRequest()
		}

		cookie, err := generator.GenerateTokenCookie(user.ID, user.Name)

		if err != nil {
			log.Println(err)
			return login.NewPostLoginInternalServerError().WithPayload(&models.Error{
				Message: "Error creating token",
			})
		}

		return login.NewPostLoginOK().WithSetCookie(cookie).WithPayload(&models.LoginSuccess{
			ID:   user.ID,
			Name: user.Name,
		})
	}
}
