//go:generate mockgen -destination=../../../mocks/adapter/commentadapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/commentadapter RepoAdapter

package commentadapter

import (
	"context"

	"golang-projects-a/pkg/core/domain"
)

type RepoAdapter interface {
	Find(ctx context.Context, id uint64) (comment domain.Comment, err error)
	GetList(ctx context.Context, filter RepoFilter) (comments domain.Comments, err error)

	Create(ctx context.Context, data RepoCreate) (comment domain.Comment, err error)
	Update(ctx context.Context, id uint64, data RepoUpdate) (comment domain.Comment, err error)
	Delete(ctx context.Context, id uint64) (err error)
}

type RepoFilter struct {
	ID             uint64
	OrganizationId uint64

	Limit  int
	Offset int
}

type RepoCreate struct {
	Comment        string
	OrganizationId uint64
}

type RepoUpdate struct {
	Comment        string
	OrganizationId uint64
}
