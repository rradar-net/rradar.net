package users

import (
	"context"

	"github.com/rradar-net/rradar.net/ent"
	"github.com/rradar-net/rradar.net/ent/user"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(ctx context.Context, user User) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client}
}

func (r repository) Create(ctx context.Context, user User) (*User, error) {
	u, err := r.client.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).Save(ctx)

	if err != nil {
		return nil, err
	}

	entity := toUserEntity(u)

	return &entity, nil
}

func (r repository) GetByUsername(ctx context.Context, username string) (*User, error) {
	u, err := r.client.User.Query().
		Where(user.Username(username)).
		Only(ctx)

	if err != nil {
		if _, ok := err.(*ent.NotSingularError); ok {
			log.Panic().Msgf("Found at least two entries with username %s. Err: %s", username, err.Error())
		}

		return nil, err
	}

	entity := toUserEntity(u)

	return &entity, nil
}
