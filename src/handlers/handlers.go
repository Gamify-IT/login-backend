// Package handlers implements the logic to handle HTTP requests.
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

// ConfigureAPI initializes our HTTP handlers and registers them with the api.
func ConfigureAPI(api *operations.LoginAPI, dbClient *db.PrismaClient) {
	jwtSecret := os.Getenv("JWT_KEY")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_KEY must not be empty"))
	}

	cookieName := os.Getenv("AUTH_COOKIE_NAME")
	if cookieName == "" {
		panic(fmt.Errorf("AUTH_COOKIE_NAME must not be empty"))
	}

	jwtValidityDuration, err := time.ParseDuration(os.Getenv("JWT_VALIDITY_DURATION"))
	if err != nil {
		panic(fmt.Errorf("could parse JWT_VALIDITY_DURATION: %w", err))
	}

	// Initialize dependencies
	generator := auth.NewAuthenticator(jwtSecret, cookieName, jwtValidityDuration)
	hasher := &hash.Bcrypt{}

	// Route: /authenticate
	api.AuthenticatePostAuthenticateHandler = AuthenticateUser(generator)

	// Route: /health
	api.HealthGetHealthHandler = healthHandler(dbClient)

	// Route: /login
	api.LoginPostLoginHandler = LoginUser(dbClient, generator, hasher)

	// Route: /register
	api.RegisterPostRegisterHandler = RegisterUser(dbClient, hasher)
}
