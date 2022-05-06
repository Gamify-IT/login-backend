package handlers

import (
	"errors"
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

// healthHandler defines the function for the endpoint /health
func healthHandler(client *db.PrismaClient) health.GetHealthHandlerFunc {
	return health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {

		// Check if DB is alive by performing any select
		_, err := client.User.FindFirst(db.User.ID.Contains("")).Exec(params.HTTPRequest.Context())
		if err != nil && !errors.Is(err, db.ErrNotFound) {
			return health.NewGetHealthServiceUnavailable().WithPayload(&models.Health{
				Status: "DOWN",
			})
		}
		return health.NewGetHealthOK().WithPayload(&models.Health{
			Status: "UP",
		})
	})
}
