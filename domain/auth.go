package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"golang.org/x/exp/slices"
)

const (
	CHARSET string = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	TOKEN_LENGTH int = 32
)

type SessionInfo struct {
	UserId      int
	TimeCreated time.Time
}

type Session map[string]SessionInfo

func (session *Session) GetSecretToken() string {
	keys := make([]string, 0, len(*session))
	for k := range *session {
		keys = append(keys, k)
	}
	return keys[0]
}

func NewSecretToken() string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)
	byteSlice := make([]byte, TOKEN_LENGTH)

	for i := range byteSlice {
		byteSlice[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}

	return string(byteSlice)
}

type AuthRepository interface {
	SignIn(userSignIn UserSignIn) (*Session, error)
	SignOut(secretToken string) error
	VerifyTokenInDb(secretToken string) error
	VerifyUserAuthorization(secretToken string, userId int) error
}

type AuthRepositoryDb struct {
	CsvDb  *clients.CsvDbClient
	JsonDb *clients.JsonDbClient
}

func NewAuthRepositoryDb(
	CsvDb *clients.CsvDbClient,
	JsonDb *clients.JsonDbClient,
) AuthRepositoryDb {
	return AuthRepositoryDb{CsvDb, JsonDb}
}

func FindUserByUsername(authRepoDb AuthRepositoryDb, userSignIn UserSignIn) (
	*User,
	error,
) {
	var users []User
	if readFileErr := authRepoDb.CsvDb.ReadCsvContents(
		&users,
		clients.FileReader{},
	); readFileErr != nil {
		return nil, readFileErr
	}

	userIdx := slices.IndexFunc(
		users,
		func(user User) bool { return user.Username == userSignIn.Username },
	)

	if userIdx < 0 {
		return nil, errors.New("user was not found")
	}

	return &users[userIdx], nil
}

func (authRepoDb AuthRepositoryDb) SignIn(userSignIn UserSignIn) (
	*Session,
	error,
) {
	foundUser, userNotFoundErr := FindUserByUsername(authRepoDb, userSignIn)

	if userNotFoundErr != nil {
		return nil, userNotFoundErr
	}

	if verifyPwdErr := VerifyPassword(
		foundUser.Password,
		userSignIn.Password,
	); verifyPwdErr != nil {
		return nil, verifyPwdErr
	}

	var authObj map[string]SessionInfo
	fileReadErr := authRepoDb.JsonDb.ReadJsonContents(&authObj, clients.FileReader{})

	if fileReadErr != nil {
		return nil, fileReadErr
	}

	secretToken := NewSecretToken()
	authObj[secretToken] = SessionInfo{
		UserId:      foundUser.Id,
		TimeCreated: time.Now(),
	}
	session := Session{
		secretToken: {
			UserId:      foundUser.Id,
			TimeCreated: time.Now(),
		},
	}

	if fileWriteErr := authRepoDb.JsonDb.UpdateJsonContents(
		authObj,
		clients.FileWriter{},
	); fileWriteErr != nil {
		return nil, fileWriteErr
	}

	return &session, nil
}

func (authRepoDb AuthRepositoryDb) SignOut(secretToken string) error {
	var authObj map[string]SessionInfo
	fileReadErr := authRepoDb.JsonDb.ReadJsonContents(&authObj, clients.FileReader{})

	if fileReadErr != nil {
		return fileReadErr
	}

	delete(authObj, secretToken)

	if fileWriteErr := authRepoDb.JsonDb.UpdateJsonContents(
		authObj,
		clients.FileWriter{},
	); fileWriteErr != nil {
		return fileWriteErr
	}

	return nil
}

func (authRepoDb AuthRepositoryDb) VerifyTokenInDb(secretToken string) error {
	var authObj map[string]SessionInfo
	fileReadErr := authRepoDb.JsonDb.ReadJsonContents(&authObj, clients.FileReader{})

	if fileReadErr != nil {
		return fileReadErr
	}

	if _, ok := authObj[secretToken]; !ok {
		return errors.New("user is not logged in")
	}

	return nil
}

func (authRepoDb AuthRepositoryDb) VerifyUserAuthorization(
	secretToken string,
	userId int,
) error {
	var authObj map[string]SessionInfo
	fileReadErr := authRepoDb.JsonDb.ReadJsonContents(&authObj, clients.FileReader{})

	if fileReadErr != nil {
		return fileReadErr
	}

	sessionInfo := authObj[secretToken]

	if sessionInfo.UserId != userId {
		return errors.New("user is not authorized")
	}

	return nil
}
