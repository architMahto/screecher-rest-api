package utils

import "net/http"

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ApiError) AsMessage() *ApiError {
	return &ApiError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
