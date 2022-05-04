package handlers

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/services"
	"github.com/architMahto/screecher-rest-api/utils"
)

type ScreechHandler struct {
	ScreechService services.ScreechService
}

func NewScreechHandler(fileDb *clients.FileDBClient) ScreechHandler {
	screechRepositoryDb := domain.NewScreechRepositoryDb(fileDb)
	screechService := services.NewScreechService(screechRepositoryDb)

	screechHandler := ScreechHandler{screechService}

	return screechHandler
}

func (screechHandler *ScreechHandler) GetAllScreeches(
	res http.ResponseWriter,
	req *http.Request,
) {
	ctxCollationConf := req.Context().Value(domain.COLLATION_CONF)
	if conf, ok := ctxCollationConf.(domain.ScreechCollationConfig); ok {
		collationConf := conf

		screeches, err := screechHandler.ScreechService.GetAllScreeches(collationConf)
		utils.WriteSuccessResponse(res, http.StatusOK, screeches)

		if err != nil {
			unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
			utils.WriteErrorResponse(res, *unexpectedErr)
			return
		}
	} else {
		unexpectedErr := utils.NewUnexpectedError("There was an unexpected error.")
		utils.WriteErrorResponse(res, *unexpectedErr)
		return
	}
}
