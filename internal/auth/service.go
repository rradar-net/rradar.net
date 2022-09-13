package auth

import (
	"github.com/rradar-net/rradar.net/internal/auth/argon2id"
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/errs"
	"github.com/rradar-net/rradar.net/internal/users"
	"github.com/rradar-net/rradar.net/pkg/proto"
	"github.com/rs/zerolog/log"
	"gopkg.in/guregu/null.v4/zero"
)

func loginUser(env env.Env, request *proto.LoginRequest) (*users.User, *errs.SentinelError) {
	user, err := env.UserRepository.GetByUsername(env.Ctx, request.Username)
	if err != nil {
		return nil, errs.NewErrUsernameOrEmailNotFound()
	}

	passwordOk, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errs.NewErrInternalServerError()
	}

	if !passwordOk {
		return nil, errs.NewErrIncorrectPassword()
	}

	return user, nil
}

func registerUser(env env.Env, request *proto.RegisterRequest) (*users.User, *errs.SentinelError) {
	// Check if username is available
	available := env.UserRepository.IsUsernameAvailable(env.Ctx, request.Username)
	if !available {
		return nil, errs.NewErrUsernameIsNotAvailable(request.Username)
	}

	// Check if email is available
	email := zero.StringFromPtr(request.Email)
	emailStr := email.ValueOrZero()
	if emailStr != "" {
		available = env.UserRepository.IsEmailAvailable(env.Ctx, emailStr)
		if !available {
			return nil, errs.NewErrEmailIsAlreadyTaken(emailStr)
		}
	}

	// Hash password
	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errs.NewErrInternalServerError()
	}
	request.Password = "" // throw away the plain password right after hashing

	// Create User entity
	user := &users.User{
		Username: request.Username,
		Password: hash,
		Email:    email,
	}

	// Save User to db
	user, err = env.UserRepository.Create(env.Ctx, *user)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errs.NewErrInternalServerError()
	}
	return user, nil
}
