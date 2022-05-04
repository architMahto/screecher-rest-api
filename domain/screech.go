package domain

import (
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
)

type Screech struct {
	Id           int       `csv:"id" json:"id"`
	Content      string    `csv:"content" json:"content"`
	CreatorId    string    `csv:"creator_id" json:"creator_id"`
	DateCreated  time.Time `csv:"date_created" json:"date_created"`
	DateModified time.Time `csv:"date_modified" json:"date_modified"`
}

type ScreechRepository interface {
	GetAllScreechesFromDB() ([]Screech, error)
}

func (screechRepoDb RepositoryDb[Screech]) GetAllScreechesFromDB() (
	[]Screech,
	error,
) {
	screeches := []Screech{}
	err := screechRepoDb.FileDB.ReadFileContents(
		&screeches,
		clients.FileReader{},
	)

	return screeches, err
}
