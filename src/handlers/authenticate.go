package handlers

import (
	"fmt"
	"github.com/Gamify-IT/login-backend/src/gen/models"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations/authenticate"
	"github.com/Gamify-IT/login-backend/src/handlers/auth"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt"
	"log"
)

// AuthenticateUser validates the "token" cookie and renews it.
func AuthenticateUser(generator *auth.Authenticator) authenticate.PostAuthenticateHandlerFunc {
	return func(params authenticate.PostAuthenticateParams) middleware.Responder {
		tokenCookie, err := params.HTTPRequest.Cookie("token")
		if err != nil {
			log.Println(fmt.Errorf("token cookie not set"))
			return authenticate.NewPostAuthenticateUnauthorized()
		}

		token, err := generator.Verify(tokenCookie.Value)
		if err != nil {
			log.Println(fmt.Errorf("verify token cookie: %w", err))
			return authenticate.NewPostAuthenticateUnauthorized()
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("error parsing claims in token cookie: %s", tokenCookie)
			return authenticate.NewPostAuthenticateUnauthorized()
		}

		if id, ok := claims["id"].(string); ok {
			if user, ok := claims["user"].(string); ok {
				newCookie, err := generator.GenerateTokenCookie(id, user)

				if err != nil {
					log.Println(fmt.Errorf("generate new token: %w", err))
					return authenticate.NewPostAuthenticateUnauthorized()
				}

				return authenticate.NewPostAuthenticateOK().WithSetCookie(newCookie).WithPayload(&models.LoginSuccess{
					ID:   id,
					Name: user,
				})
			}
		}

		log.Printf("id or user not set in token cookie: %s", tokenCookie)
		return authenticate.NewPostAuthenticateUnauthorized()
	}
}
