package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/domain"
	"time"
)

type organizationRepo struct {
	col *mongo.Collection
}

var _ organizationadapter.RepoAdapter = organizationRepo{}

func (o organizationRepo) Find(ctx context.Context, id uint64) (domain.Organization, error) {
	var res Organization
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	if err := o.col.FindOne(ctx, selector).Decode(&res); err != nil {
		return res.domain(), err
	}

	return res.domain(), nil
}

func (o organizationRepo) GetList(ctx context.Context, filter organizationadapter.RepoFilter) (domain.Organizations, error) {
	var res domain.Organizations

	selector, option, err := o.buildSelectorFind(filter)
	if err != nil {
		return res, err
	}

	cur, err := o.col.Find(ctx, selector, option)
	if err != nil {
		return res, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var organization Organization
		err := cur.Decode(&organization)
		if err != nil {
			return res, err
		}

		res.Items = append(res.Items, organization.domain())
		res.RowCount++
	}

	return res, nil
}

func (o organizationRepo) buildSelectorFind(filter organizationadapter.RepoFilter) (map[string]interface{}, *options.FindOptions, error) {
	selector := make(map[string]interface{})
	option := options.Find()

	BuildSelectorUint64(selector, "id", filter.ID)
	BuildSelectorString(selector, "name", filter.Name)
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	option.Limit = GetLimit(filter.Limit)
	option.Skip = GetSkip(filter.Offset)

	return selector, option, nil
}

func (o organizationRepo) Create(ctx context.Context, data organizationadapter.RepoCreate) (domain.Organization, error) {
	newId, err := GetId(ctx, o.col)
	if err != nil {
		return domain.Organization{}, err
	}
	organization := Organization{
		ID:        newId,
		Name:      data.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = o.col.InsertOne(ctx, organization)
	if err != nil {
		return domain.Organization{}, err
	}

	newOrganization, err := o.Find(ctx, newId)
	if err != nil {
		return domain.Organization{}, err
	}

	return newOrganization, nil
}

func (o organizationRepo) Update(ctx context.Context, id uint64, data organizationadapter.RepoUpdate) (domain.Organization, error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	organization := Organization{
		Name:      data.Name,
		UpdatedAt: time.Now(),
	}

	result, err := o.col.UpdateOne(ctx, selector, bson.M{"$set": organization})
	if err != nil {
		return domain.Organization{}, err
	}

	if result.MatchedCount < 1 {
		return domain.Organization{}, adapter.ErrNotFound
	}

	newOrganization, err := o.Find(ctx, id)
	if err != nil {
		return domain.Organization{}, err
	}

	return newOrganization, nil
}

func (o organizationRepo) Delete(ctx context.Context, id uint64) (err error) {
	selector := make(map[string]interface{})
	selector["id"] = id
	selector["deleted_at"] = GetSoftDeletedSelector(false)
	organization := Organization{
		DeletedAt: time.Now(),
	}

	result, err := o.col.UpdateOne(ctx, selector, bson.M{"$set": organization})
	if err != nil {
		return err
	}

	if result.MatchedCount < 1 {
		return adapter.ErrNotFound
	}

	return nil
}

func (o organizationRepo) FindByName(ctx context.Context, name string) (organization domain.Organization, err error) {
	var res Organization
	selector := make(map[string]interface{})
	selector["name"] = name
	selector["deleted_at"] = GetSoftDeletedSelector(false)

	if err := o.col.FindOne(ctx, selector).Decode(&res); err != nil {
		return res.domain(), err
	}

	return res.domain(), nil
}
