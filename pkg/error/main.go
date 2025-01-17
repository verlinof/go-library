package pkg_error

import (
	"net/http"
)

type ClientError struct {
	Status  int
	Message string
}

func NewBadRequest(err error) *ClientError {
	return &ClientError{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func NewInternalServerError(err error) *ClientError {
	return &ClientError{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func NewNotFound(err error) *ClientError {
	return &ClientError{
		Status:  http.StatusNotFound,
		Message: err.Error(),
	}
}

func NewForbidden(err error) *ClientError {
	return &ClientError{
		Status:  http.StatusForbidden,
		Message: err.Error(),
	}
}

func NewUnauthorized(err error) *ClientError {
	return &ClientError{
		Status:  http.StatusUnauthorized,
		Message: err.Error(),
	}
}
