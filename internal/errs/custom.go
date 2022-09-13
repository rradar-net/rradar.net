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

var ErrEmailIsAlreadyTaken = errors.New("email is already taken")

func NewErrEmailIsAlreadyTaken(email string) *SentinelError {
	return &SentinelError{
		http.StatusConflict,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data: map[string]string{
				"email": fmt.Sprintf("Email %s is already taken.", email),
			},
		},
		ErrEmailIsAlreadyTaken,
	}
}

var ErrUsernameOrEmailNotFound = errors.New("username or email not found")

func NewErrUsernameOrEmailNotFound() *SentinelError {
	return &SentinelError{
		http.StatusNotFound,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data: map[string]string{
				"username": "Couldn't find that username or email.",
			},
		},
		ErrUsernameOrEmailNotFound,
	}
}

var ErrIncorrectPassword = errors.New("incorrect password")

func NewErrIncorrectPassword() *SentinelError {
	return &SentinelError{
		http.StatusForbidden,
		proto.ErrorResponse{
			Status: proto.Status_fail,
			Data: map[string]string{
				"password": "Incorrect password.",
			},
		},
		ErrIncorrectPassword,
	}
}
