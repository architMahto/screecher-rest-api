package domain

import (
	"github.com/architMahto/screecher-rest-api/app/clients"
)

type DataModel interface {
	User | Screech
}

type RepositoryDb[T DataModel] struct {
	FileDB *clients.FileDBClient
}
