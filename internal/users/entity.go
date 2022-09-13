package users

import (
	"github.com/rradar-net/rradar.net/ent"
	"gopkg.in/guregu/null.v4/zero"
)

type User struct {
	Username string
	Password string
	Email    zero.String
}

func toUserEntity(u *ent.User) User {
	return User{
		Username: u.Username,
		Password: u.Password,
		Email:    zero.StringFromPtr(u.Email),
	}
}
