//go:generate mockgen -destination=../../../mocks/adapter/rolepermissionadapter/repoadapter.go github.com/alchasni/golang-projects-a/pkg/core/adapter/rolepermissionadapter RepoAdapter

package rolepermissionadapter

import (
	"context"

	"golang-projects-a/pkg/types"
)

type RepoAdapter interface {
	CheckExistence(ctx context.Context, roleCode types.Code, permCode types.Code) (err error)
}
