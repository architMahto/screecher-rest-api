package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteErrorResponse(res http.ResponseWriter, err ApiError) {
	res.Header().Add("Content-Type", "application/json")

	statusCode := err.Code

	res.WriteHeader(statusCode)
	encodingErr := json.NewEncoder(res).Encode(err)

	if encodingErr != nil {
		fmt.Fprintf(res, "%s", encodingErr.Error())
	}
}

func WriteSuccessResponse(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)

	if err != nil {
		fmt.Fprintf(res, "%s", err.Error())
	}
}
