package domain

import (
	"errors"
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"golang.org/x/exp/slices"
)

type User struct {
	Id              int       `csv:"id" json:"id"`
	FirstName       string    `csv:"first_name" json:"first_name"`
	LastName        string    `csv:"last_name" json:"last_name"`
	Username        string    `csv:"username" json:"username"`
	Password        string    `csv:"password" json:"-"`
	SecretToken     string    `csv:"secret_token" json:"-"`
	ProfileImageURI string    `csv:"profile_image" json:"profile_image_uri"`
	DateCreated     time.Time `csv:"date_created" json:"date_created"`
	DateModified    time.Time `csv:"date_modified" json:"date_modified"`
}

type UserRepository interface {
	GetAllUsersFromDB() ([]User, error)
	GetUserFromDb(userId int) (*User, error)
}

type UserRepositoryDb struct {
	FileDB *clients.FileDBClient
}

func NewUserRepositoryDb(FileDB *clients.FileDBClient) UserRepositoryDb {
	return UserRepositoryDb{FileDB}
}

func (userRepoDb UserRepositoryDb) GetAllUsersFromDB() (
	[]User,
	error,
) {
	users := []User{}
	err := userRepoDb.FileDB.ReadFileContents(&users, clients.FileReader{})

	return users, err
}

func (userRepoDb UserRepositoryDb) GetUserFromDb(userId int) (
	*User,
	error,
) {
	users := []User{}
	if readFileErr := userRepoDb.FileDB.ReadFileContents(
		&users,
		clients.FileReader{},
	); readFileErr != nil {
		return nil, readFileErr
	}

	userIdx := slices.IndexFunc(
		users,
		func(user User) bool { return user.Id == userId },
	)

	if userIdx < 0 {
		return nil, errors.New("User was not found")
	}

	return &users[userIdx], nil
}
