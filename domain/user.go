package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

type User struct {
	Id              int       `csv:"id" json:"id"`
	FirstName       string    `csv:"first_name" json:"first_name"`
	LastName        string    `csv:"last_name" json:"last_name"`
	Username        string    `csv:"username" json:"username"`
	Password        string    `csv:"password" json:"password"`
	ProfileImageURI string    `csv:"profile_image" json:"profile_image_uri"`
	DateCreated     time.Time `csv:"date_created" json:"date_created"`
	DateModified    time.Time `csv:"date_modified" json:"date_modified"`
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (usr User) MarshalJSON() ([]byte, error) {
	var tmpUser struct {
		Id              int       `csv:"id" json:"id"`
		FirstName       string    `csv:"first_name" json:"first_name"`
		LastName        string    `csv:"last_name" json:"last_name"`
		Username        string    `csv:"username" json:"username"`
		ProfileImageURI string    `csv:"profile_image" json:"profile_image_uri"`
		DateCreated     time.Time `csv:"date_created" json:"date_created"`
		DateModified    time.Time `csv:"date_modified" json:"date_modified"`
	}
	tmpUser.Id = usr.Id
	tmpUser.FirstName = usr.FirstName
	tmpUser.LastName = usr.LastName
	tmpUser.Username = usr.Username
	tmpUser.ProfileImageURI = usr.ProfileImageURI
	tmpUser.DateCreated = usr.DateCreated
	tmpUser.DateModified = usr.DateModified

	return json.Marshal(&tmpUser)
}

func (user *User) PrepareForAddition() error {
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Id = 0
	user.DateCreated = time.Now()
	user.DateModified = time.Now()
	user.Password = string(hashedPassword)

	return nil
}

func (user *User) DoesUsernameExist(users []User) bool {
	for _, u := range users {
		if u.Username == user.Username {
			return true
		}
	}
	return false
}

func (user *User) Validate() error {
	if len(user.Username) == 0 ||
		len(user.FirstName) == 0 ||
		len(user.LastName) == 0 ||
		len(user.Password) == 0 {
		return errors.New("please enter a username, password, first name, and last name")
	}

	if len(user.Username) > 80 {
		return errors.New("username length is too long")
	}
	if len(user.FirstName) > 100 {
		return errors.New("first name length is too long")
	}
	if len(user.LastName) > 80 {
		return errors.New("last name length is too long")
	}

	return nil
}

type UserRepository interface {
	GetAllUsersFromDB() ([]User, error)
	GetUserFromDb(userId int) (*User, error)
	AddUserToDB(user *User) (*User, error)
}

type UserRepositoryDb struct {
	FileDB *clients.FileDBClient
}

func NewUserRepositoryDb(FileDB *clients.FileDBClient) UserRepositoryDb {
	return UserRepositoryDb{FileDB}
}

func FetchAllUsersFromDB(userRepoDb UserRepositoryDb) (
	[]User,
	error,
) {
	users := []User{}
	if readFileErr := userRepoDb.FileDB.ReadFileContents(
		&users,
		clients.FileReader{},
	); readFileErr != nil {
		return nil, readFileErr
	}

	return users, nil
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

func (userRepoDb UserRepositoryDb) AddUserToDB(user *User) (
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

	usersListSize := len(users)

	if usersListSize > 0 {
		lastUser := users[len(users)-1]
		user.Id = lastUser.Id + 1
	} else {
		user.Id = 1
	}

	users = append(users, *user)

	if writeFileErr := userRepoDb.FileDB.UpdateFileContents(
		&users,
		clients.FileWriter{},
	); writeFileErr != nil {
		return nil, writeFileErr
	}

	return user, nil
}
