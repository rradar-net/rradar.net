package errs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rradar-net/rradar.net/pkg/proto"
)

var ErrValidation = errors.New("validation error")

func NewErrValidation(data map[string]string) *SentinelError {
	return &SentinelError{
		http.StatusBadRequest,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data:   data,
		},
		ErrValidation,
	}
}

var ErrInternalServerError = errors.New("internal server error")

func NewErrInternalServerError() *SentinelError {
	msg := "Internal server error."
	return &SentinelError{
		http.StatusInternalServerError,
		proto.ErrorResponse{
			Status:  proto.Status_error,
			Message: &msg,
		},
		ErrInternalServerError,
	}
}

var ErrUsernameIsNotAvailable = errors.New("user already exists")

func NewErrUsernameIsNotAvailable(username string) *SentinelError {
	return &SentinelError{
		http.StatusConflict,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data: map[string]string{
				"username": fmt.Sprintf("Username %s is not available.", username),
			},
		},
		ErrUsernameIsNotAvailable,
	}
}

func NewErrEmailIsAlreadyTaken(email string) *SentinelError {
	return &SentinelError{
		http.StatusConflict,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data: map[string]string{
				"email": fmt.Sprintf("Email %s is already taken.", email),
			},
		},
		ErrUsernameIsNotAvailable,
	}
}
