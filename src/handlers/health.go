package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

func healthHandler() health.GetHealthHandlerFunc {
	return health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
		return health.NewGetHealthOK().WithPayload(&models.Health{
			Components: &models.HealthComponents{
				Database: "UP",
			},
			Status: "UP",
		})
	})
}
