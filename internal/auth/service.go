package auth

import (
	"github.com/rradar-net/rradar.net/internal/auth/argon2id"
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/errs"
	"github.com/rradar-net/rradar.net/internal/users"
	"github.com/rradar-net/rradar.net/pkg/proto"
	"github.com/rs/zerolog/log"
)

func registerUser(env env.Env, request *proto.RegisterRequest) (*users.User, *errs.SentinelError) {
	// Check if username is available
	_, err := env.UserRepository.GetByUsername(env.Ctx, request.Username)
	if err == nil {
		return nil, errs.NewErrUsernameIsNotAvailable(request.Username)
	}

	// Hash password
	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errs.NewErrInternalServerError()
	}
	request.Password = ""

	// Create User entity
	user := &users.User{
		Username: request.Username,
		Password: hash,
		Email:    proto.OptionalString(request.Email),
	}

	// Save User to db
	user, err = env.UserRepository.Create(env.Ctx, *user)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errs.NewErrInternalServerError()
	}

	return user, nil
}
