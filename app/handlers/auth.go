package handlers

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/services"
	"github.com/architMahto/screecher-rest-api/utils"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(csvDb *clients.CsvDbClient, jsonDb *clients.JsonDbClient) AuthHandler {
	authRepositoryDb := domain.NewAuthRepositoryDb(csvDb, jsonDb)
	authService := services.NewAuthService(authRepositoryDb)

	authHandler := AuthHandler{authService}

	return authHandler
}

func (authHandler *AuthHandler) SignIn(
	res http.ResponseWriter,
	req *http.Request,
) {
	ctxUser := req.Context().Value(domain.COLLATION_CONF)
	userSignIn, ok := ctxUser.(domain.UserSignIn)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	session, err := authHandler.AuthService.SignIn(userSignIn)

	if err != nil {
		unexpectedErr := utils.NewAuthenticationError("Could not login")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, session)
}

func (authHandler *AuthHandler) SignOut(
	res http.ResponseWriter,
	req *http.Request,
) {
	authHeader := req.Header.Get("Authorization")

	err := authHandler.AuthService.SignOut(authHeader)

	if err != nil {
		unexpectedErr := utils.NewAuthenticationError("User is not authenticated")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, "User is logged out")
}
