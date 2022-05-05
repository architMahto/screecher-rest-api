package middleware

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/architMahto/screecher-rest-api/domain"
	"github.com/architMahto/screecher-rest-api/utils"
)

func ValidateNumQueryParam(numQueryParamStr string) (*int, error) {
	value, err := strconv.Atoi(numQueryParamStr)

	if err != nil {
		return nil, err
	}

	return &value, nil
}

func ValidateScreechQueryParams(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		pageNumStr := req.URL.Query().Get("pageNum")
		pageSizeStr := req.URL.Query().Get("pageSize")
		sortStr := req.URL.Query().Get("sort")

		collationConf := domain.ScreechCollationConfig{
			PageNum:      1,
			PageSize:     domain.MIN_PAGE_SIZE,
			SortOrderDir: domain.DESC_SORT_ORDER,
		}

		if pageNumStr != "" {
			pageNum, pageNumReadErr := ValidateNumQueryParam(pageNumStr)

			if pageNumReadErr != nil {
				pageNumValidationErr := utils.NewValidationError("pageNum should be a number")
				utils.WriteErrorResponse(res, *pageNumValidationErr)
				return
			}

			collationConf.PageNum = *pageNum
		}

		if pageSizeStr != "" {
			pageSize, pageSizeReadErr := ValidateNumQueryParam(pageSizeStr)

			if pageSizeReadErr != nil ||
				*pageSize < domain.MIN_PAGE_SIZE ||
				*pageSize > domain.MAX_PAGE_SIZE ||
				*pageSize%50 != 0 {
				errMessage := "pageSize should be a number between 50 to 500 that is divisible by 50"
				pageNumValidationErr := utils.NewValidationError(errMessage)
				utils.WriteErrorResponse(res, *pageNumValidationErr)
				return
			}

			collationConf.PageSize = *pageSize
		}

		if sortStr == "" {
			collationConf.SortOrderDir = domain.DESC_SORT_ORDER
		} else if sortStr == domain.DESC_SORT_ORDER || sortStr == domain.ASC_SORT_ORDER {
			collationConf.SortOrderDir = sortStr
		} else {
			pageNumValidationErr := utils.NewValidationError("sort should be either asc or desc")
			utils.WriteErrorResponse(res, *pageNumValidationErr)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), domain.COLLATION_CONF, collationConf))
		next.ServeHTTP(res, req)
	})
}

func ValidateScreechBody(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		reqBody, readErr := ioutil.ReadAll(req.Body)

		if readErr != nil {
			validationError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *validationError)
			return
		}

		screech := domain.Screech{}

		if unmarshalErr := json.Unmarshal(reqBody, &screech); unmarshalErr != nil {
			unmarshalError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *unmarshalError)
			return
		}

		screech.PrepareForAddition()

		if validationErr := screech.Validate(); validationErr != nil {
			validationError := utils.NewValidationError(validationErr.Error())
			utils.WriteErrorResponse(res, *validationError)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), domain.COLLATION_CONF, screech))
		next.ServeHTTP(res, req)
	})
}

func ValidateUserCreateReqBody(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		reqBody, readErr := ioutil.ReadAll(req.Body)

		if readErr != nil {
			validationError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *validationError)
			return
		}

		user := domain.User{}

		if unmarshalErr := json.Unmarshal(reqBody, &user); unmarshalErr != nil {
			unmarshalError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *unmarshalError)
			return
		}

		if validationErr := user.ValidateFields(); validationErr != nil {
			validationErr := utils.NewValidationError(validationErr.Error())
			utils.WriteErrorResponse(res, *validationErr)
			return
		}

		if passwordHashErr := user.PrepareForAddition(); passwordHashErr != nil {
			passwordHashErr := utils.NewUnexpectedError("Unexpected error")
			utils.WriteErrorResponse(res, *passwordHashErr)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), domain.COLLATION_CONF, user))
		next.ServeHTTP(res, req)
	})
}

func ValidateUserUpdateReqBody(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		reqBody, readErr := ioutil.ReadAll(req.Body)

		if readErr != nil {
			validationError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *validationError)
			return
		}

		userUpdateBody := domain.UserUpdateBody{}

		if unmarshalErr := json.Unmarshal(reqBody, &userUpdateBody); unmarshalErr != nil {
			unmarshalError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *unmarshalError)
			return
		}

		if validationErr := userUpdateBody.ValidateFields(); validationErr != nil {
			validationErr := utils.NewValidationError(validationErr.Error())
			utils.WriteErrorResponse(res, *validationErr)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), domain.COLLATION_CONF, userUpdateBody))
		next.ServeHTTP(res, req)
	})
}

func ValidateUserSignInReqBody(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		reqBody, readErr := ioutil.ReadAll(req.Body)

		if readErr != nil {
			validationError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *validationError)
			return
		}

		userSignIn := domain.UserSignIn{}

		if unmarshalErr := json.Unmarshal(reqBody, &userSignIn); unmarshalErr != nil {
			unmarshalError := utils.NewValidationError("Issues with input format")
			utils.WriteErrorResponse(res, *unmarshalError)
			return
		}

		if validationErr := userSignIn.ValidateFields(); validationErr != nil {
			validationErr := utils.NewValidationError(validationErr.Error())
			utils.WriteErrorResponse(res, *validationErr)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), domain.COLLATION_CONF, userSignIn))
		next.ServeHTTP(res, req)
	})
}
