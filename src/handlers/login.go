package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/login"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

// loginUser let a user log in
func loginUser(api *operations.LoginAPI, client *db.PrismaClient) login.PostLoginHandlerFunc {
	return login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
		username := params.Body.Username
		password := params.Body.Password

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

		user, err := client.User.FindFirst(db.User.Name.Equals(*username)).Exec(params.HTTPRequest.Context())

		if err != nil {
			return login.NewPostLoginInternalServerError()
		}

		if user == nil {
			return login.NewPostLoginBadRequest()
		}

		if user.PasswordHash != string(hashedPassword) {
			return login.NewPostLoginBadRequest()
		}

		return login.NewPostLoginOK()

	})
}
