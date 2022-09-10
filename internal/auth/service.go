package auth

import (
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/errors"
	"github.com/rradar-net/rradar.net/internal/users"
	"github.com/rradar-net/rradar.net/pkg/proto"
)

func registerUser(env env.Env, request *proto.RegisterRequest) (*users.User, *errors.SentinelError) {
	user := &users.User{
		Username: request.Username,
		Password: request.Password,
		Email:    proto.OptionalString(request.Email),
	}

	user, err := env.UserRepository.CreateUser(env.Ctx, *user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
