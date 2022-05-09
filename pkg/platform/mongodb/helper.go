package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-projects-a/pkg/core/adapter"
)

type IdTracker struct {
	Id  string `bson:"_id"`
	Seq uint64 `bson:"seq"`
}

func BuildSelectorString(s map[string]interface{}, field string, value string) map[string]interface{} {
	selector := s
	if value != "" {
		selector[field] = value
	}
	return selector
}

func BuildSelectorUint64(s map[string]interface{}, field string, value uint64) map[string]interface{} {
	selector := s
	if value != 0 {
		selector[field] = value
	}
	return selector
}

func GetFindOption(limit int64, skip int64) *options.FindOptions {
	option := options.Find()
	option.Limit = getLimit(limit)
	option.Skip = getSkip(skip)

	return option

}

func getLimit(limit int64) *int64 {
	if limit == 0 {
		limit = DefaultLimit
	} else if limit > MaxLimit {
		limit = MaxLimit
	}
	return &limit
}

func getSkip(skip int64) *int64 {
	return &skip
}

func GetSoftDeletedSelector(deleted bool) map[string]interface{} {
	return map[string]interface{}{"$exists": deleted}
}

func GetId(ctx context.Context, col *mongo.Collection) (uint64, error) {
	idTracker := IdTracker{}
	res := col.FindOne(ctx, bson.M{"id": col.Name()})
	err := res.Decode(&idTracker)
	if err != nil || idTracker.Seq == 0 {
		_, err = col.InsertOne(ctx, bson.M{"id": col.Name(), "seq": 1})
		if err != nil {
			return 0, err
		}
		return 1, nil
	} else {
		filter := bson.M{"id": col.Name()}
		update := bson.M{"$inc": bson.M{"seq": 1}}
		result, err := col.UpdateOne(ctx, filter, update)
		if err != nil {
			return 0, err
		} else if result.MatchedCount == 1 {
			err := col.FindOne(ctx, bson.M{"id": col.Name()}).Decode(&idTracker)
			if err != nil {
				return 0, err
			}
			return idTracker.Seq, nil
		} else {
			return 0, adapter.ErrGenerateId
		}
	}
}
