package services

import (
	"sort"

	"github.com/architMahto/screecher-rest-api/domain"
)

type ScreechService interface {
	GetAllScreeches(collationConf domain.ScreechCollationConfig) (
		[]domain.Screech,
		error,
	)
}

type ScreechServiceHandler struct {
	ScreechRepo domain.RepositoryDb[domain.Screech]
}

func (service ScreechServiceHandler) GetAllScreeches(
	collationConf domain.ScreechCollationConfig,
) (
	[]domain.Screech,
	error,
) {
	screeches, err := service.ScreechRepo.GetAllScreechesFromDB()

	if err != nil {
		return nil, err
	}

	sort.Slice(screeches, func(i, j int) bool {
		if collationConf.SortOrderDir == domain.ASC_SORT_ORDER {
			return screeches[i].DateCreated.Before(screeches[j].DateCreated)
		}
		return screeches[i].DateCreated.After(screeches[j].DateCreated)
	})

	start, end := domain.GetPaginatedScreechesIndices(screeches, collationConf)

	return screeches[start:end], err
}

func NewScreechService(
	repo domain.RepositoryDb[domain.Screech],
) ScreechServiceHandler {
	return ScreechServiceHandler{repo}
}
