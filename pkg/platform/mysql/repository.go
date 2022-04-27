package mysql

import (
	"golang-projects-a/pkg/core/adapter/permissionadapter"
	"golang-projects-a/pkg/core/adapter/roleadapter"
	"golang-projects-a/pkg/core/adapter/rolepermissionadapter"

	"gorm.io/gorm"
)

const (
	table_Roles           = "roles"
	table_RolePermissions = "role_permissions"
	table_Permissions     = "permissions"
)

type Service struct {
	db        *gorm.DB
	paginator paginator
}

func New(cfg Config) (Service, error) {
	db, err := initORM(cfg.DB)
	if err != nil {
		return Service{}, err
	}

	return Service{
		db: db,
		paginator: paginator{
			minPageSize: cfg.Pagination.MinPageSize,
			maxPageSize: cfg.Pagination.MaxPageSize,
		},
	}, nil
}

func (s Service) RoleRepo() roleadapter.RepoAdapter {
	return roleRepo{
		db:        s.db,
		paginator: s.paginator,
	}
}

func (s Service) PermissionRepo() permissionadapter.RepoAdapter {
	return permissionRepo{
		db:        s.db,
		paginator: s.paginator,
	}
}

func (s Service) RolePermissionRepo() rolepermissionadapter.RepoAdapter {
	return rolePermissionRepo{
		db:        s.db,
		paginator: s.paginator,
	}
}
