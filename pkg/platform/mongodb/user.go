package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"
	"log"
	"time"
)

type userRepo struct {
	col *mongo.Collection
}

var _ useradapter.RepoAdapter = userRepo{}

func (u userRepo) Find(ctx context.Context, id uint64) (domain.User, error) {
	var res User
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	if err := u.col.FindOne(ctx, selector).Decode(&res); err != nil {
		log.Fatal(err)
	}

	return res.domain(), nil
}

func (u userRepo) GetList(ctx context.Context, filter useradapter.RepoFilter) (domain.Users, error) {
	var res domain.Users

	selector, option, err := u.buildSelectorFind(&filter)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := u.col.Find(ctx, selector, option)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err.Error())
		}

		res.Items = append(res.Items, user.domain())
		res.RowCount++
	}

	return res, nil
}

func (u userRepo) buildSelectorFind(filter *useradapter.RepoFilter) (map[string]interface{}, *options.FindOptions, error) {
	selector := make(map[string]interface{})
	option := options.Find()

	if filter == nil {
		return selector, option, fmt.Errorf("empty search form for overall daily cumulative")
	}
	if filter.ID == 0 {
		return selector, option, fmt.Errorf("empty ID")
	}

	selector["id"] = filter.ID
	selector["username"] = filter.Username
	selector["email"] = filter.Email
	selector["password"] = filter.Password
	selector["avatar_url"] = filter.AvatarUrl

	option.Limit = GetLimit(filter.Limit)
	option.Skip = GetSkip(filter.Offset)

	return selector, option, nil
}

func (u userRepo) Create(ctx context.Context, data useradapter.RepoCreate) (domain.User, error) {
	newId, err := GetId(ctx, u.col)
	if err != nil {
		log.Fatal(err)
	}
	user := User{
		ID:        uint64(newId),
		Username:  data.Username,
		Password:  data.Password,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = u.col.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	newUser, err := u.Find(ctx, uint64(newId))
	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil
}

func (u userRepo) Update(ctx context.Context, id uint64, data useradapter.RepoUpdate) (domain.User, error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	user := User{
		Username:  data.Username,
		Email:     data.Email,
		UpdatedAt: time.Now(),
	}

	result, err := u.col.UpdateOne(ctx, selector, bson.M{"$set": user})
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount < 1 {
		log.Fatal(err)
	}

	newUser, err := u.Find(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil
}

func (u userRepo) Delete(ctx context.Context, id uint64) (err error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	user := User{
		DeletedAt: time.Now(),
	}

	result, err := u.col.UpdateOne(ctx, selector, bson.M{"$set": user})
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount < 1 {
		log.Fatal(err)
	}

	return nil
}
