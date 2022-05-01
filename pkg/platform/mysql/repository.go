package mysql

import (
	"golang-projects-a/pkg/core/adapter/useradapter"
	"gorm.io/gorm"
)

const (
	tableUsers = "users"
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

func (s Service) UserRepo() useradapter.RepoAdapter {
	return userRepo{
		db:        s.db,
		paginator: s.paginator,
	}
}
