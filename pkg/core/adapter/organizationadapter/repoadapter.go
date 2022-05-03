//go:generate mockgen -destination=../../../mocks/adapter/organizationadapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/organizationadapter RepoAdapter

package organizationadapter

import (
	"context"

	"golang-projects-a/pkg/core/domain"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint64) (organization domain.Organization, err error)
	GetList(ctx context.Context, filter RepoFilter) (organizations domain.Organizations, err error)

	Create(ctx context.Context, data RepoCreate) (organization domain.Organization, err error)
	Update(ctx context.Context, id uint64, data RepoUpdate) (organization domain.Organization, err error)
	Delete(ctx context.Context, id uint64) (err error)

	FindByName(ctx context.Context, name string) (organization domain.Organization, err error)
}

type RepoFilter struct {
	ID   uint64
	Name string

	Limit  int
	Offset int
}

type RepoCreate struct {
	Name string
}

type RepoUpdate struct {
	Name string
}
