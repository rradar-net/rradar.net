package users

import "github.com/rradar-net/rradar.net/ent"

type User struct {
	Username string
	Password string
	Email    string
}

func toUserEntity(u *ent.User) User {
	return User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}
