package handlers

import (
	"net/http"
	"strconv"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/services"
	"github.com/architMahto/screecher-rest-api/utils"
	"github.com/gorilla/mux"
)

type ScreechHandler struct {
	ScreechService services.ScreechService
}

func NewScreechHandler(csvDb *clients.CsvDbClient) ScreechHandler {
	screechRepositoryDb := domain.NewScreechRepositoryDb(csvDb)
	screechService := services.NewScreechService(screechRepositoryDb)

	screechHandler := ScreechHandler{screechService}

	return screechHandler
}

func (screechHandler *ScreechHandler) GetScreeches(
	res http.ResponseWriter,
	req *http.Request,
) {
	ctxCollationConf := req.Context().Value(domain.COLLATION_CONF)
	collationConf, ok := ctxCollationConf.(domain.ScreechCollationConfig)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	screeches, err := screechHandler.ScreechService.GetAllScreeches(collationConf)

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, screeches)
}

func (screechHandler *ScreechHandler) GetScreechById(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	screechId, _ := strconv.Atoi(vars["screech_id"])

	user, err := screechHandler.ScreechService.GetScreechById(screechId)

	if err != nil && user == nil {
		notFoundErr := utils.NewNotFoundError("Screech was not found.")
		utils.WriteErrorResponse(res, *notFoundErr)
		return
	}

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, user)
}

func (screechHandler *ScreechHandler) CreateScreech(
	res http.ResponseWriter,
	req *http.Request,
) {
	ctxScreech := req.Context().Value(domain.COLLATION_CONF)
	screech, ok := ctxScreech.(domain.Screech)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	screechResult, err := screechHandler.ScreechService.CreateNewScreech(&screech)

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, screechResult)
}

func (screechHandler *ScreechHandler) UpdateScreech(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	screechId, _ := strconv.Atoi(vars["screech_id"])
	ctxScreech := req.Context().Value(domain.COLLATION_CONF)
	screech, ok := ctxScreech.(domain.Screech)

	if !ok {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	screech.PrepareForUpdate(screechId)
	screechResult, err := screechHandler.ScreechService.UpdateScreech(&screech)

	if err != nil {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}

	utils.WriteSuccessResponse(res, http.StatusOK, screechResult)
}
