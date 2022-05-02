package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/register"
	"github.com/go-openapi/runtime/middleware"
)

func registerUser(api *operations.LoginAPI, client *db.PrismaClient) register.PostRegisterHandlerFunc {
	return register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
		username := params.Username
		email := params.Email
		password := params.Password

		// TODO: Database Check
		success := (username == username) && (password == password) && (email == email)
		if success {
			return register.NewPostRegisterOK()
		} else {
			return register.NewPostRegisterBadRequest()
		}

	})
}
