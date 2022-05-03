//go:generate mockgen -destination=../../../mocks/adapter/useradapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/useradapter RepoAdapter

package useradapter

import (
	"context"

	"golang-projects-a/pkg/core/domain"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint64) (user domain.User, err error)
	GetList(ctx context.Context, filter RepoFilter) (users domain.Users, err error)

	Create(ctx context.Context, data RepoCreate) (user domain.User, err error)
	Update(ctx context.Context, id uint64, data RepoUpdate) (user domain.User, err error)
	Delete(ctx context.Context, id uint64) (err error)
}

type RepoFilter struct {
	ID             uint64
	Username       string
	Email          string
	Password       string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64

	Limit  int64
	Offset int64
}

type RepoCreate struct {
	Username       string
	Email          string
	Password       string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
}

type RepoUpdate struct {
	Username       string
	Email          string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
}
