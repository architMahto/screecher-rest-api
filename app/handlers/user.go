package handlers

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/services"
	"github.com/architMahto/screecher-rest-api/utils"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(fileDb *clients.FileDBClient) UserHandler {
	userRepositoryDb := domain.NewUserRepositoryDb(fileDb)
	userService := services.NewUserService(userRepositoryDb)

	userHandler := UserHandler{userService}

	return userHandler
}

func (userHandler *UserHandler) GetAllUsers(
	res http.ResponseWriter,
	req *http.Request,
) {
	users, err := userHandler.UserService.GetAllUsers()

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}
	utils.WriteSuccessResponse(res, http.StatusOK, users)
}
