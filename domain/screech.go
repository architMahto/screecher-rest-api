package domain

import (
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
)

type key string

const (
	COLLATION_CONF key = "collationConf"
)

type Screech struct {
	Id           int       `csv:"id" json:"id"`
	Content      string    `csv:"content" json:"content"`
	CreatorId    string    `csv:"creator_id" json:"creator_id"`
	DateCreated  time.Time `csv:"date_created" json:"date_created"`
	DateModified time.Time `csv:"date_modified" json:"date_modified"`
}

type ScreechCollationConfig struct {
	PageNum      int
	PageSize     int
	SortOrderDir string
}

const (
	MIN_PAGE_SIZE   int    = 50
	MAX_PAGE_SIZE   int    = 500
	ASC_SORT_ORDER  string = "asc"
	DESC_SORT_ORDER string = "desc"
)

type ScreechRepository interface {
	GetAllScreechesFromDB() ([]Screech, error)
}

func GetPaginatedScreechesIndices(
	screeches []Screech,
	collationConf ScreechCollationConfig,
) (int, int) {
	screechesLen := len(screeches)
	start := (collationConf.PageNum - 1) * collationConf.PageSize

	if start > screechesLen {
		start = screechesLen
	}

	end := start + collationConf.PageSize

	if end > screechesLen {
		end = screechesLen
	}

	return start, end
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
