package handlers

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/utils"
)

func HandleNotFound(res http.ResponseWriter, req *http.Request) {
	notFoundErr := utils.NewNotFoundError("Route does not exist")
	utils.WriteErrorResponse(res, *notFoundErr)
}
