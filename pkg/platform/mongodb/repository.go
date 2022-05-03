package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"os"
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

func (s Service) OrganizationRepo() organizationadapter.RepoAdapter {
	collection := s.db.Collection(OrganizationCollection)
	s.createOrganizationIndex(collection)
	return organizationRepo{
		col: collection,
	}
}

func (s Service) UserRepo() useradapter.RepoAdapter {
	collection := s.db.Collection(UserCollection)
	s.createUserIndex(collection)
	return userRepo{
		col: collection,
	}
}

func (s Service) createOrganizationIndex(col *mongo.Collection) {
	ctx := context.Background()
	mod := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := col.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println("create index organization ERROR:", err)
		os.Exit(1)
	}
}

func (s Service) createUserIndex(col *mongo.Collection) {
	ctx := context.Background()
	mod := mongo.IndexModel{
		Keys: bson.M{
			"organization_id": 1,
		},
		Options: nil,
	}
	_, err := col.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println("create index user ERROR:", err)
		os.Exit(1)
	}
}
