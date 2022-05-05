package domain

import (
	"errors"
	"sort"
	"time"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"golang.org/x/exp/slices"
)

type key string

const (
	COLLATION_CONF key = "collationConf"
)

type Screech struct {
	Id           int       `csv:"id" json:"id"`
	Content      string    `csv:"content" json:"content"`
	CreatorId    int       `csv:"creator_id" json:"creator_id"`
	DateCreated  time.Time `csv:"date_created" json:"date_created"`
	DateModified time.Time `csv:"date_modified" json:"date_modified"`
}

func (screech *Screech) PrepareForAddition() {
	screech.Id = 0
	screech.DateCreated = time.Now()
	screech.DateModified = time.Now()
}

func (screech *Screech) PrepareForUpdate(screechId int) {
	screech.Id = screechId
	screech.DateModified = time.Now()
}

func (screech *Screech) Validate() error {
	if len(screech.Content) > 1024 {
		return errors.New("Screech content is too long")
	}

	return nil
}

type ScreechCollationConfig struct {
	PageNum      int
	PageSize     int
	SortOrderDir string
}

const (
	MIN_PAGE_SIZE      int    = 50
	MAX_PAGE_SIZE      int    = 500
	ASC_SORT_ORDER     string = "asc"
	DESC_SORT_ORDER    string = "desc"
	MAX_SCREECH_LENGTH int    = 1024
)

type ScreechRepository interface {
	GetScreechesFromDB(collationConf ScreechCollationConfig) ([]Screech, error)
	GetScreechFromDb(screechId int) (*Screech, error)
	AddScreechToDB(screech *Screech) (*Screech, error)
	UpdateScreechInDB(screech *Screech) (*Screech, error)
}

type ScreechRepositoryDb struct {
	CsvDb *clients.CsvDbClient
}

func NewScreechRepositoryDb(CsvDb *clients.CsvDbClient) ScreechRepositoryDb {
	return ScreechRepositoryDb{CsvDb}
}

func FetchAllScreechesFromDB(screechRepoDb ScreechRepositoryDb) (
	[]Screech,
	error,
) {
	screeches := []Screech{}
	if readFileErr := screechRepoDb.CsvDb.ReadCsvContents(
		&screeches,
		clients.FileReader{},
	); readFileErr != nil {
		return nil, readFileErr
	}

	return screeches, nil
}

func FindScreechById(screeches []Screech, screechId int) (
	*int,
	error,
) {
	screechIdx := slices.IndexFunc(
		screeches,
		func(screech Screech) bool { return screech.Id == screechId },
	)

	if screechIdx < 0 {
		return nil, errors.New("Screech was not found")
	}

	return &screechIdx, nil
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

func (screechRepoDb ScreechRepositoryDb) GetScreechesFromDB(
	collationConf ScreechCollationConfig,
) (
	[]Screech,
	error,
) {
	screeches, readFileErr := FetchAllScreechesFromDB(screechRepoDb)

	if readFileErr != nil {
		return nil, readFileErr
	}

	sort.Slice(screeches, func(i, j int) bool {
		if collationConf.SortOrderDir == ASC_SORT_ORDER {
			return screeches[i].DateCreated.Before(screeches[j].DateCreated)
		}
		return screeches[i].DateCreated.After(screeches[j].DateCreated)
	})

	start, end := GetPaginatedScreechesIndices(screeches, collationConf)

	return screeches[start:end], nil
}

func (screechRepoDb ScreechRepositoryDb) GetScreechFromDb(
	screechId int,
) (
	*Screech,
	error,
) {
	screeches, readFileErr := FetchAllScreechesFromDB(screechRepoDb)

	if readFileErr != nil {
		return nil, readFileErr
	}

	screechIdx, notFoundErr := FindScreechById(screeches, screechId)

	return &screeches[*screechIdx], notFoundErr
}

func (screechRepoDb ScreechRepositoryDb) AddScreechToDB(
	screech *Screech,
) (
	*Screech,
	error,
) {
	screeches, readFileErr := FetchAllScreechesFromDB(screechRepoDb)

	if readFileErr != nil {
		return nil, readFileErr
	}

	lastScreech := screeches[len(screeches)-1]
	screech.Id = lastScreech.Id + 1

	screeches = append(screeches, *screech)

	if writeFileErr := screechRepoDb.CsvDb.UpdateCsvContents(
		&screeches,
		clients.FileWriter{},
	); writeFileErr != nil {
		return nil, writeFileErr
	}

	return screech, nil
}

func (screechRepoDb ScreechRepositoryDb) UpdateScreechInDB(
	screech *Screech,
) (
	*Screech,
	error,
) {
	screeches, readFileErr := FetchAllScreechesFromDB(screechRepoDb)

	if readFileErr != nil {
		return nil, readFileErr
	}

	screechIdx, notFoundErr := FindScreechById(screeches, screech.Id)
	updatedScreech := Screech{
		Id:           screech.Id,
		Content:      screech.Content,
		CreatorId:    screech.CreatorId,
		DateCreated:  screeches[*screechIdx].DateCreated,
		DateModified: time.Now(),
	}
	screeches[*screechIdx] = updatedScreech

	if writeFileErr := screechRepoDb.CsvDb.UpdateCsvContents(
		&screeches,
		clients.FileWriter{},
	); writeFileErr != nil {
		return nil, writeFileErr
	}

	return &updatedScreech, notFoundErr
}
