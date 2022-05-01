package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
)

type Service struct {
	db *mongo.Database
}

func New(cfg Config) (Service, error) {
	database, err := initDB(cfg.DB)
	if err != nil {
		return Service{}, err
	}

	return Service{
		db: database,
	}, nil
}

func (s Service) UserRepo() useradapter.RepoAdapter {
	collection := s.db.Collection(UserCollection)
	return userRepo{
		col: collection,
	}
}

func (s Service) OrganizationRepo() organizationadapter.RepoAdapter {
	collection := s.db.Collection(OrganizationCollection)
	return organizationRepo{
		col: collection,
	}
}
