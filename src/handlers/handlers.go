package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/todos"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureAPI(api *operations.LoginAPI, client *db.PrismaClient) {
	// Route: /health
	api.HealthGetHealthHandler = healthHandler(client)

	// Route: /login
	api.LoginPostLoginHandler = loginUser(api, client)

	// Route: /register
	api.RegisterPostRegisterHandler = registerUser(api, client)
}
