package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
}

type UserServiceHandler struct {
	UserRepo domain.RepositoryDb[domain.User]
}

func (service UserServiceHandler) GetAllUsers() (
	[]domain.User,
	error,
) {
	users, err := service.UserRepo.GetAllUsersFromDB()

	if err != nil {
		return nil, err
	}

	return users, err
}

func NewUserService(
	repo domain.RepositoryDb[domain.User],
) UserServiceHandler {
	return UserServiceHandler{repo}
}
