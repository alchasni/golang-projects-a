package mysql

import (
	"context"

	"twatter/pkg/core/adapter/rolepermissionadapter"
	"twatter/pkg/types"

	"gorm.io/gorm"
)

type rolePermissionRepo struct {
	db        *gorm.DB
	paginator paginator
}

var _ rolepermissionadapter.RepoAdapter = rolePermissionRepo{}

func (r rolePermissionRepo) CheckExistence(ctx context.Context, roleCode types.Code, permCode types.Code) (err error) {
	panic("implement me")
}
