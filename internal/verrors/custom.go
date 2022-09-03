package verrors

import (
	"net/http"

	"github.com/rradar-net/rradar.net/pkg/proto"
)

type SentinelError struct {
	HttpStatus int
	Error      proto.ErrorResponse
}

func (s *SentinelError) JSON() *proto.ErrorResponse {
	return &s.Error
}

func ErrValidation(data map[string]string) *SentinelError {
	return &SentinelError{
		http.StatusBadRequest,
		proto.ErrorResponse{
			Status: proto.Status_Fail,
			Data:   data,
		},
	}
}

func ErrInternalServerError() *SentinelError {
	msg := "internal server error"
	return &SentinelError{
		http.StatusInternalServerError,
		proto.ErrorResponse{
			Status:  proto.Status_Error,
			Message: &msg,
		},
	}
}
