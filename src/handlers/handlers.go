package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"os"
)

func ConfigureAPI(api *operations.LoginAPI, client *db.PrismaClient) {
	generator := auth.NewAuthenticator(os.Getenv("JWT_KEY"))
	hasher := &hash.Bcrypt{}

	// Route: /health
	api.HealthGetHealthHandler = healthHandler(client)

	// Route: /login
	api.LoginPostLoginHandler = LoginUser(client, generator, hasher)

	// Route: /register
	api.RegisterPostRegisterHandler = RegisterUser(client, hasher)
}
