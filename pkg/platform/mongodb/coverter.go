package mongodb

import (
	"golang-projects-a/pkg/core/domain"
	"time"
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
		ID:             user.ID,
		Username:       user.Username,
		Password:       user.Password,
		Email:          user.Email,
		AvatarUrl:      user.AvatarUrl,
		OrganizationId: user.OrganizationId,
		FollowingCount: user.FollowingCount,
		FollowerCount:  user.FollowerCount,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      time.Time{},
	}
}

func (u User) domain() domain.User {
	return domain.User{
		ID:             u.ID,
		Username:       u.Username,
		Password:       u.Password,
		Email:          u.Email,
		AvatarUrl:      u.AvatarUrl,
		OrganizationId: u.OrganizationId,
		FollowingCount: u.FollowingCount,
		FollowerCount:  u.FollowerCount,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		DeletedAt:      time.Time{},
	}
}
