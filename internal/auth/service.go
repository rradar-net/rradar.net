package auth

import (
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/users"
	"github.com/rradar-net/rradar.net/internal/verrors"
	"github.com/rradar-net/rradar.net/pkg/proto"
)

func registerUser(env env.Env, request *proto.RegisterRequest) (*users.User, *verrors.SentinelError) {
	user := users.User{
		Username: request.Username,
		Password: request.Password,
		Email:    proto.OptionalString(request.Email),
	}

	result := env.Db.Create(&user)
	if result.Error != nil {
		return nil, verrors.ErrInternalServerError()
	}

	return &user, nil
}
