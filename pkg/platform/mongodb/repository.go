package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
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
	collection := s.db.Collection("user")
	return userRepo{
		col: collection,
	}
}
