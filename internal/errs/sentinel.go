package errs

import "github.com/rradar-net/rradar.net/pkg/proto"

type SentinelError struct {
	HttpStatus int
	ErrProto   proto.ErrorResponse
	origin     error
}

func (s *SentinelError) Unwrap() error {
	return s.origin
}

func (s *SentinelError) Error() string {
	return *s.ErrProto.Message
}

func (s *SentinelError) JSON() *proto.ErrorResponse {
	return &s.ErrProto
}
