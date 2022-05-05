package middleware

import (
	"errors"
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/handlers"
	"github.com/architMahto/screecher-rest-api/utils"
)

func GetAuthHeaderVal(res http.ResponseWriter, req *http.Request) (*string, error) {
	authHeader := req.Header.Get("Authorization")

	if authHeader == "" {
		authError := utils.NewAuthenticationError("User is not logged in")
		utils.WriteErrorResponse(res, *authError)
		return nil, errors.New("user is not logged in")
	}

	return &authHeader, nil
}

func IsUserAuthenticated(
	authHandler handlers.AuthHandler,
	next http.HandlerFunc,
) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader, authHeaderEmptyErr := GetAuthHeaderVal(res, req)

		if authHeaderEmptyErr != nil {
			return
		}

		tokenNotFoundErr := authHandler.AuthService.VerifyTokenInDb(*authHeader)

		if tokenNotFoundErr != nil {
			authError := utils.NewAuthenticationError("User is not logged in")
			utils.WriteErrorResponse(res, *authError)
			return
		}

		next.ServeHTTP(res, req)
	})
}
