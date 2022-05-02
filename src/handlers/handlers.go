package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
)

func ConfigureAPI(api *operations.LoginAPI, client *db.PrismaClient) {
	// Route: /health
	api.HealthGetHealthHandler = healthHandler(client)
}
