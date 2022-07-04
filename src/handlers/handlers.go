package handlers

import (
	"fmt"
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/Gamify-IT/login-backend/src/handlers/hash"
	"os"
	"time"
)

// ConfigureAPI registers all our HTTP handlers with the Swagger API server.
func ConfigureAPI(api *operations.LoginAPI, client *db.PrismaClient) {
	jwtSecret := os.Getenv("JWT_KEY")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_KEY must not be empty"))
	}

	jwtValidityDuration, err := time.ParseDuration(os.Getenv("JWT_VALIDITY_DURATION"))
	if err != nil {
		panic(fmt.Errorf("could parse JWT_VALIDITY_DURATION: %e", err))
	}

	generator := auth.NewAuthenticator(jwtSecret, jwtValidityDuration)
	hasher := &hash.Bcrypt{}

	// Route: /health
	api.HealthGetHealthHandler = healthHandler(client)

	// Route: /login
	api.LoginPostLoginHandler = LoginUser(client, generator, hasher)

	// Route: /register
	api.RegisterPostRegisterHandler = RegisterUser(client, hasher)
}
