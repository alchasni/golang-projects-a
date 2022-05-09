package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/domain"
	"time"
)

type commentRepo struct {
	col *mongo.Collection
}

var _ commentadapter.RepoAdapter = commentRepo{}

func (c commentRepo) Find(ctx context.Context, id uint64) (domain.Comment, error) {
	var res Comment
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	if err := c.col.FindOne(ctx, selector).Decode(&res); err != nil {
		return res.domain(), err
	}

	return res.domain(), nil
}

func (c commentRepo) GetList(ctx context.Context, filter commentadapter.RepoFilter) (domain.Comments, error) {
	var res domain.Comments

	selector, err := c.buildSelectorFind(filter)
	if err != nil {
		return res, err
	}
	option := GetFindOption(filter.Limit, filter.Offset)

	cur, err := c.col.Find(ctx, selector, option)
	if err != nil {
		return res, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var comment Comment
		err := cur.Decode(&comment)
		if err != nil {
			return res, err
		}

		res.Items = append(res.Items, comment.domain())
		res.RowCount++
	}

	return res, nil
}

func (c commentRepo) buildSelectorFind(filter commentadapter.RepoFilter) (map[string]interface{}, error) {
	selector := make(map[string]interface{})

	BuildSelectorUint64(selector, "id", filter.ID)
	BuildSelectorUint64(selector, "organization_id", filter.OrganizationId)
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	return selector, nil
}

func (c commentRepo) Create(ctx context.Context, data commentadapter.RepoCreate) (domain.Comment, error) {
	newId, err := GetId(ctx, c.col)
	if err != nil {
		return domain.Comment{}, err
	}
	comment := Comment{
		ID:             newId,
		Comment:        data.Comment,
		OrganizationId: data.OrganizationId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err = c.col.InsertOne(ctx, comment)
	if err != nil {
		return domain.Comment{}, err
	}

	newComment, err := c.Find(ctx, newId)
	if err != nil {
		return domain.Comment{}, err
	}

	return newComment, nil
}

func (c commentRepo) Update(ctx context.Context, id uint64, data commentadapter.RepoUpdate) (domain.Comment, error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	comment := Comment{
		Comment:        data.Comment,
		OrganizationId: data.OrganizationId,
		UpdatedAt:      time.Now(),
	}

	result, err := c.col.UpdateOne(ctx, selector, bson.M{"$set": comment})
	if err != nil {
		return domain.Comment{}, err
	}

	if result.MatchedCount < 1 {
		return domain.Comment{}, adapter.ErrNotFound
	}

	newComment, err := c.Find(ctx, id)
	if err != nil {
		return domain.Comment{}, err
	}

	return newComment, nil
}

func (c commentRepo) Delete(ctx context.Context, id uint64) (err error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	comment := Comment{
		DeletedAt: time.Now(),
	}

	result, err := c.col.UpdateOne(ctx, selector, bson.M{"$set": comment})
	if err != nil {
		return err
	}

	if result.MatchedCount < 1 {
		return adapter.ErrNotFound
	}

	return nil
}

func (c commentRepo) DeleteMany(ctx context.Context, filter commentadapter.RepoFilter) (err error) {
	selector, err := c.buildSelectorFind(filter)
	if err != nil {
		return err
	}

	comment := Comment{
		DeletedAt: time.Now(),
	}

	result, err := c.col.UpdateMany(ctx, selector, bson.M{"$set": comment})
	if err != nil {
		return err
	}

	if result.MatchedCount < 1 {
		return adapter.ErrNotFound
	}

	return nil
}
