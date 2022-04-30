package handlers

import (
	"github.com/Gamify-IT/login-backend/src/gen/db"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/todos"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureAPI(api *operations.LoginAPI, client *db.PrismaClient) {
	if api.TodosGetHandler == nil {
		api.TodosGetHandler = todos.GetHandlerFunc(func(params todos.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.Get has not yet been implemented")
		})
	}
}
