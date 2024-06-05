package custom_errors

import (
	"errors"
	"fmt"
	"net/http"

	"to-do-list-bts.id/constants"
)

var (
	ErrInvalidAuthToken = errors.New(constants.InvalidAuthTokenErrMsg)
)

type AppError struct {
	Code    int
	Message string
	err     error
}

func (e AppError) Error() string {
	return fmt.Sprint(e.Message)
}

func BadRequest(err error, message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		err:     err,
	}
}

func NotFound(err error) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: constants.ResponseMsgErrorNotFound,
		err:     err,
	}
}

func Unauthorized(err error, message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
		err:     err,
	}
}

func InvalidAuthToken() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: constants.InvalidAuthTokenErrMsg,
		err:     ErrInvalidAuthToken,
	}
}
