package users

import (
	"context"

	"github.com/rradar-net/rradar.net/ent"
	"github.com/rradar-net/rradar.net/ent/user"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(ctx context.Context, user User) (*User, error)
	IsUsernameAvailable(ctx context.Context, username string) bool
	IsEmailAvailable(ctx context.Context, email string) bool
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
		SetNillableEmail(user.Email.Ptr()).Save(ctx)

	if err != nil {
		return nil, err
	}

	entity := toUserEntity(u)

	return &entity, nil
}

func (r repository) IsUsernameAvailable(ctx context.Context, username string) bool {
	count, err := r.client.User.Query().
		Where(user.Username(username)).
		Count(ctx)

	if err != nil {
		log.Panic().Msg(err.Error())
	}

	return count == 0
}

func (r repository) IsEmailAvailable(ctx context.Context, email string) bool {
	count, err := r.client.User.Query().
		Where(user.Email(email)).
		Count(ctx)

	if err != nil {
		log.Panic().Msg(err.Error())
	}

	return count == 0
}
