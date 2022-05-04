package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type ScreechService interface {
	GetAllScreeches(collationConf domain.ScreechCollationConfig) (
		[]domain.Screech,
		error,
	)
	GetScreechById(screechId int) (*domain.Screech, error)
}

type ScreechServiceHandler struct {
	ScreechRepo domain.ScreechRepositoryDb
}

func (service ScreechServiceHandler) GetAllScreeches(
	collationConf domain.ScreechCollationConfig,
) (
	[]domain.Screech,
	error,
) {
	screeches, err := service.ScreechRepo.GetAllScreechesFromDB(
		collationConf,
	)

	if err != nil {
		return nil, err
	}

	return screeches, err
}

func (service ScreechServiceHandler) GetScreechById(screechId int) (
	*domain.Screech,
	error,
) {
	screech, err := service.ScreechRepo.GetScreechFromDb(screechId)

	if err != nil {
		return nil, err
	}

	return screech, err
}

func NewScreechService(
	repo domain.ScreechRepositoryDb,
) ScreechServiceHandler {
	return ScreechServiceHandler{repo}
}
