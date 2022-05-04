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
	CreateNewScreech(screech *domain.Screech) (*domain.Screech, error)
	UpdateScreech(screech *domain.Screech) (*domain.Screech, error)
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
	screeches, err := service.ScreechRepo.GetScreechesFromDB(
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

func (service ScreechServiceHandler) CreateNewScreech(screech *domain.Screech) (
	*domain.Screech,
	error,
) {
	screechRes, err := service.ScreechRepo.AddScreechToDB(screech)

	if err != nil {
		return nil, err
	}

	return screechRes, err
}

func (service ScreechServiceHandler) UpdateScreech(screech *domain.Screech) (
	*domain.Screech,
	error,
) {
	screechRes, err := service.ScreechRepo.UpdateScreechInDB(screech)

	if err != nil {
		return nil, err
	}

	return screechRes, err
}

func NewScreechService(
	repo domain.ScreechRepositoryDb,
) ScreechServiceHandler {
	return ScreechServiceHandler{repo}
}
