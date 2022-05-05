package handlers

import (
	"net/http"
	"strconv"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/services"
	"github.com/architMahto/screecher-rest-api/utils"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(csvDb *clients.CsvDbClient) UserHandler {
	userRepositoryDb := domain.NewUserRepositoryDb(csvDb)
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

func (userHandler *UserHandler) GetUserById(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	userId, _ := strconv.Atoi(vars["user_id"])

	user, err := userHandler.UserService.GetUserById(userId)

	if err != nil && user == nil {
		notFoundErr := utils.NewNotFoundError("User was not found.")
		utils.WriteErrorResponse(res, *notFoundErr)
		return
	}

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, user)
}

func (userHandler *UserHandler) CreateUser(
	res http.ResponseWriter,
	req *http.Request,
) {
	ctxUser := req.Context().Value(domain.COLLATION_CONF)
	user, ok := ctxUser.(domain.User)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	users, fetchUsersErr := userHandler.UserService.GetAllUsers()

	if fetchUsersErr != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	if userExistsErr := user.DoesUsernameExist(users); userExistsErr {
		userExistsErr := utils.NewConflictError("User with username already exists.")
		utils.WriteErrorResponse(res, *userExistsErr)
		return
	}

	userResult, err := userHandler.UserService.CreateNewUser(&user)

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, userResult)
}

func (userHandler *UserHandler) UpdateUser(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	userId, _ := strconv.Atoi(vars["user_id"])
	ctxUser := req.Context().Value(domain.COLLATION_CONF)
	userUpdateBody, ok := ctxUser.(domain.UserUpdateBody)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	// userUpdateBody["Id"] = userId
	userResult, err := userHandler.UserService.UpdateUser(userId, userUpdateBody)

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, userResult)
}
