package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type AuthService interface {
	SignIn(userSignIn domain.UserSignIn) (*domain.Session, error)
	SignOut(secretToken string) error
	VerifyTokenInDb(secretToken string) error
	VerifyUserAuthorization(secretToken string, userId int) error
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

func (service AuthServiceHandler) VerifyUserAuthorization(
	secretToken string,
	userId int,
) error {
	err := service.AuthRepo.VerifyUserAuthorization(secretToken, userId)

	if err != nil {
		return err
	}

	return nil
}

func (service AuthServiceHandler) VerifyTokenInDb(secretToken string) error {
	err := service.AuthRepo.VerifyTokenInDb(secretToken)

	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo domain.AuthRepositoryDb) AuthServiceHandler {
	return AuthServiceHandler{repo}
}
