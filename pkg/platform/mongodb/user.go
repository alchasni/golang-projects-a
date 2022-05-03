package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"
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
		return res.domain(), err
	}

	return res.domain(), nil
}

func (u userRepo) GetList(ctx context.Context, filter useradapter.RepoFilter) (domain.Users, error) {
	var res domain.Users

	selector, option, err := u.buildSelectorFind(filter)
	if err != nil {
		return res, err
	}

	cur, err := u.col.Find(ctx, selector, option)
	if err != nil {
		return res, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return res, err
		}

		res.Items = append(res.Items, user.domain())
		res.RowCount++
	}

	return res, nil
}

func (u userRepo) buildSelectorFind(filter useradapter.RepoFilter) (map[string]interface{}, *options.FindOptions, error) {
	selector := make(map[string]interface{})
	option := options.Find()

	BuildSelectorUint64(selector, "id", filter.ID)
	BuildSelectorString(selector, "username", filter.Username)
	BuildSelectorString(selector, "email", filter.Email)
	BuildSelectorString(selector, "password", filter.Password)
	BuildSelectorString(selector, "avatar_url", filter.AvatarUrl)
	BuildSelectorUint64(selector, "organization_id", filter.OrganizationId)
	BuildSelectorUint64(selector, "following_count", filter.FollowingCount)
	BuildSelectorUint64(selector, "follower_count", filter.FollowerCount)
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	option.Limit = GetLimit(filter.Limit)
	option.Skip = GetSkip(filter.Offset)

	return selector, option, nil
}

func (u userRepo) Create(ctx context.Context, data useradapter.RepoCreate) (domain.User, error) {
	newId, err := GetId(ctx, u.col)
	if err != nil {
		return domain.User{}, err
	}
	user := User{
		ID:             newId,
		Username:       data.Username,
		Password:       data.Password,
		Email:          data.Email,
		AvatarUrl:      data.AvatarUrl,
		OrganizationId: data.OrganizationId,
		FollowingCount: data.FollowingCount,
		FollowerCount:  data.FollowerCount,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      time.Time{},
	}

	_, err = u.col.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	newUser, err := u.Find(ctx, newId)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func (u userRepo) Update(ctx context.Context, id uint64, data useradapter.RepoUpdate) (domain.User, error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	user := User{
		Username:       data.Username,
		Email:          data.Email,
		OrganizationId: data.OrganizationId,
		FollowingCount: data.FollowingCount,
		FollowerCount:  data.FollowerCount,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Now(),
	}

	result, err := u.col.UpdateOne(ctx, selector, bson.M{"$set": user})
	if err != nil {
		return domain.User{}, err
	}

	if result.MatchedCount < 1 {
		return domain.User{}, adapter.ErrNotFound
	}

	newUser, err := u.Find(ctx, id)
	if err != nil {
		return domain.User{}, err
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
		return err
	}

	if result.MatchedCount < 1 {
		return adapter.ErrNotFound
	}

	return nil
}
