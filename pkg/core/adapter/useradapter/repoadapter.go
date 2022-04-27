//go:generate mockgen -destination=../../../mocks/adapter/useradapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/useradapter RepoAdapter

package useradapter

import (
	"context"

	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/types"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint64) (user domain.User, err error)
	GetList(ctx context.Context, filter RepoFilter) (users domain.Users, err error)

	Create(ctx context.Context, data RepoCreate) (user domain.User, err error)
	Update(ctx context.Context, id uint64, data RepoUpdate) (user domain.User, err error)
	Delete(ctx context.Context, id uint64) (err error)

	Validate(ctx context.Context, username string, password string) (err error)
	UpdatePassword(ctx context.Context, username string, newPassword string) (err error)
}

type RepoFilter struct {
	Username string
	Email    string
	RoleCode types.Code
	PageNo   int
	PageSize int
}

type RepoCreate struct {
	Username string
	Email    string
	Password string
	RoleCode types.Code
}

type RepoUpdate struct {
	Email    string
	RoleCode types.Code
}
