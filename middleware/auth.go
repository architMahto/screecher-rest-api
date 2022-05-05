package middleware

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/utils"
)

func DoesAuthHeaderExist(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		if authHeader == "" {
			authError := utils.NewAuthenticationError("User is not logged in")
			utils.WriteErrorResponse(res, *authError)
			return
		}

		next.ServeHTTP(res, req)
	})
}
