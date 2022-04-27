//go:generate mockgen -destination=../../../mocks/adapter/rolepermissionadapter/repoadapter.go github.com/adityarev/go-be-starter-2/pkg/core/adapter/rolepermissionadapter RepoAdapter

package rolepermissionadapter

import (
	"context"

	"twatter/pkg/types"
)

type RepoAdapter interface {
	CheckExistence(ctx context.Context, roleCode types.Code, permCode types.Code) (err error)
}
