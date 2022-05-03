package handlers

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/utils"
)

func Home(res http.ResponseWriter, req *http.Request) {
	utils.WriteSuccessResponse(
		res,
		http.StatusOK,
		"Welcome to the Screecher REST API",
	)
}
