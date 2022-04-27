//go:generate mockgen -destination=../../../mocks/adapter/permissionadapter/repoadapter.go github.com/adityarev/go-be-starter-2/pkg/core/adapter/permissionadapter RepoAdapter

package permissionadapter

import (
	"context"

	"twatter/pkg/core/domain"
	"twatter/pkg/types"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint32) (permission domain.Permission, err error)
	GetList(ctx context.Context, filter RepoFilter) (permissions domain.Permissions, err error)

	Create(ctx context.Context, data RepoCreate) (permission domain.Permission, err error)
	Update(ctx context.Context, id uint32, data RepoUpdate) (permission domain.Permission, err error)
	Delete(ctx context.Context, id uint32) (err error)
}

type RepoFilter struct {
	Code       types.Code
	Name       string
	Restricted string
	PageNo     int
	PageSize   int
}

type RepoCreate struct {
	Code       types.Code
	Name       string
	Restricted string
}

type RepoUpdate struct {
	Name       string
	Restricted string
}
