package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IdTracker struct {
	Id  string `bson:"_id"`
	Seq uint   `bson:"seq"`
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

func GetLimit(l int) *int64 {
	limit := int64(l)
	if limit == 0 {
		limit = DefaultLimit
	} else if limit > MaxLimit {
		limit = MaxLimit
	}
	return &limit
}

func GetSkip(s int) *int64 {
	skip := int64(s)
	return &skip
}

func GetSoftDeletedSelector(deleted bool) map[string]interface{} {
	return map[string]interface{}{"$exists": deleted}
}

func GetId(ctx context.Context, col *mongo.Collection) (uint, error) {
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
			return 0, errors.New("generate id failed")
		}
	}
}
