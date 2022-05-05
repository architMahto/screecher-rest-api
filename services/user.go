package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(userId int) (*domain.User, error)
	CreateNewUser(user *domain.User) (*domain.User, error)
	UpdateUser(
		userId int,
		userUpdateBody domain.UserUpdateBody,
	) (*domain.User, error)
}

type UserServiceHandler struct {
	UserRepo domain.UserRepositoryDb
}

func (service UserServiceHandler) GetAllUsers() (
	[]domain.User,
	error,
) {
	users, err := service.UserRepo.GetUsersFromDB()

	if err != nil {
		return nil, err
	}

	return users, err
}

func (service UserServiceHandler) GetUserById(userId int) (
	*domain.User,
	error,
) {
	user, err := service.UserRepo.GetUserFromDb(userId)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (service UserServiceHandler) CreateNewUser(user *domain.User) (
	*domain.User,
	error,
) {
	userRes, err := service.UserRepo.AddUserToDB(user)

	if err != nil {
		return nil, err
	}

	return userRes, err
}

func (service UserServiceHandler) UpdateUser(
	userId int,
	userUpdateBody domain.UserUpdateBody,
) (
	*domain.User,
	error,
) {
	userRes, err := service.UserRepo.UpdateUserInDB(userId, userUpdateBody)

	if err != nil {
		return nil, err
	}

	return userRes, err
}

func NewUserService(
	repo domain.UserRepositoryDb,
) UserServiceHandler {
	return UserServiceHandler{repo}
}
