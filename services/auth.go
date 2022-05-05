package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type AuthService interface {
	SignIn(userSignIn domain.UserSignIn) (*domain.Session, error)
	SignOut(secretToken string) error
}

type AuthServiceHandler struct {
	AuthRepo domain.AuthRepositoryDb
}

func (service AuthServiceHandler) SignIn(userSignIn domain.UserSignIn) (
	*domain.Session,
	error,
) {
	session, err := service.AuthRepo.SignIn(userSignIn)

	if err != nil {
		return nil, err
	}

	return session, err
}

func (service AuthServiceHandler) SignOut(secretToken string) error {
	err := service.AuthRepo.SignOut(secretToken)

	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo domain.AuthRepositoryDb) AuthServiceHandler {
	return AuthServiceHandler{repo}
}
