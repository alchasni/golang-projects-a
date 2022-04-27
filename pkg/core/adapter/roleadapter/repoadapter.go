//go:generate mockgen -destination=../../../mocks/adapter/roleadapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/roleadapter RepoAdapter

package roleadapter

import (
	"context"

	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/types"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint32) (role domain.Role, err error)
	GetList(ctx context.Context, filter RepoFilter) (roles domain.Roles, err error)

	Create(ctx context.Context, data RepoCreate) (role domain.Role, err error)
	Update(ctx context.Context, id uint32, data RepoUpdate) (role domain.Role, err error)
	Delete(ctx context.Context, id uint32) (err error)
}

type RepoFilter struct {
	Code            types.Code
	Name            string
	PermissionCodes []types.Code
	PageNo          int
	PageSize        int
}

func (r RepoFilter) OmitemptyPermissionCodes() []types.Code {
	permCodes := make([]types.Code, 0)
	for _, permCode := range r.PermissionCodes {
		if !permCode.Empty() {
			permCodes = append(permCodes, permCode)
		}
	}
	return permCodes
}

type RepoCreate struct {
	Code            types.Code
	Name            string
	PermissionCodes []types.Code
}

type RepoUpdate struct {
	Name            string
	PermissionCodes []types.Code
}
