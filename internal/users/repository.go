package users

import (
	"context"

	"github.com/rradar-net/rradar.net/ent"
	"github.com/rradar-net/rradar.net/internal/errors"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) (*User, *errors.SentinelError)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client}
}

func (r repository) CreateUser(ctx context.Context, user User) (*User, *errors.SentinelError) {
	u, err := r.client.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).Save(ctx)

	if err != nil {
		return nil, errors.ErrInternalServerError()
	}

	entity := toUserEntity(u)

	return &entity, nil
}
