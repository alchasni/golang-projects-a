package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	User string `yaml:"user" validate:"required"`
	Pass string `yaml:"pass" validate:"required" logger:"-"`
	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`
	DB   string `yaml:"db" validate:"required"`
}

func (cfg DBConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
}

func initDB(cfg DBConfig) (*mongo.Database, error) {
	ctx := context.TODO()
	// TODO: set pass user and DB schema
	clientOptions := options.Client().ApplyURI(cfg.GetURI())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client.Database(cfg.DB), nil
}
