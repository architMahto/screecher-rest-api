package middleware

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/handlers"
	"github.com/architMahto/screecher-rest-api/utils"
)

func IsUserAuthenticated(
	authHandler handlers.AuthHandler,
	next http.HandlerFunc,
) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		if authHeader == "" {
			authError := utils.NewAuthenticationError("User is not logged in")
			utils.WriteErrorResponse(res, *authError)
			return
		}

		tokenNotFoundErr := authHandler.AuthService.VerifyTokenInDb(authHeader)

		if tokenNotFoundErr != nil {
			authError := utils.NewAuthenticationError("User is not logged in")
			utils.WriteErrorResponse(res, *authError)
			return
		}

		next.ServeHTTP(res, req)
	})
}
