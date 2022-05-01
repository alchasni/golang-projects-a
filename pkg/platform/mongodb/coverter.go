package mongodb

import (
	"golang-projects-a/pkg/core/domain"
)

func (u User) load(user domain.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u User) domain() domain.User {
	return domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		AvatarUrl: u.AvatarUrl,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
