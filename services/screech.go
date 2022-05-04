package services

import (
	"github.com/architMahto/screecher-rest-api/domain"
)

type ScreechService interface {
	GetAllScreeches(collationConf domain.ScreechCollationConfig) (
		[]domain.Screech,
		error,
	)
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

func NewScreechService(
	repo domain.ScreechRepositoryDb,
) ScreechServiceHandler {
	return ScreechServiceHandler{repo}
}
