package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter/commentadapter"
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

func (s Service) CommentRepo() commentadapter.RepoAdapter {
	collection := s.db.Collection(CommentCollection)
	s.createCommentIndex(collection)
	return commentRepo{
		col: collection,
	}
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

func (s Service) createCommentIndex(col *mongo.Collection) {
	ctx := context.Background()
	idIndex := mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		},
		Options: nil,
	}
	orgIdIndex := mongo.IndexModel{
		Keys: bson.M{
			"organization_id": 1,
		},
		Options: nil,
	}
	indexModels := []mongo.IndexModel{idIndex, orgIdIndex}
	_, err := col.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		fmt.Println("create index user ERROR:", err)
		os.Exit(1)
	}
}

func (s Service) createOrganizationIndex(col *mongo.Collection) {
	ctx := context.Background()
	idIndex := mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		},
		Options: nil,
	}
	nameIndex := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	indexModels := []mongo.IndexModel{idIndex, nameIndex}
	_, err := col.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		fmt.Println("create index organization ERROR:", err)
		os.Exit(1)
	}
}

func (s Service) createUserIndex(col *mongo.Collection) {
	ctx := context.Background()
	idIndex := mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		},
		Options: nil,
	}
	orgIdIndex := mongo.IndexModel{
		Keys: bson.M{
			"organization_id": 1,
		},
		Options: nil,
	}
	indexModels := []mongo.IndexModel{idIndex, orgIdIndex}
	_, err := col.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		fmt.Println("create index user ERROR:", err)
		os.Exit(1)
	}
}

func (s Service) Drop(ctx context.Context) {
	err := s.db.Drop(ctx)
	if err != nil {
		fmt.Println("drop DB ERROR:", err)
		os.Exit(1)
	}
}
