package mongodb

import (
	"golang-projects-a/pkg/core/domain"
)

func (o Organization) load(organization domain.Organization) Organization {
	return Organization{
		ID:        organization.ID,
		Name:      organization.Name,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}
}

func (o Organization) domain() domain.Organization {
	return domain.Organization{
		ID:        o.ID,
		Name:      o.Name,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

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
