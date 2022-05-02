package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/login"
	"github.com/go-openapi/runtime/middleware"
)

func loginUser(api *operations.LoginAPI, client *db.PrismaClient) login.PostLoginHandlerFunc {
	return login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
		username := params.Body.Username
		password := params.Body.Password

		// TODO: Database Check
		success := (username == username) && (password == password)
		if success {
			return login.NewPostLoginOK()
		} else {
			return login.NewPostLoginForbidden()
		}

	})
}
