package repository

import (
	"context"

	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	InsertOne(ctx context.Context, collectionName string, document interface{}) error
	Find(ctx context.Context, collectionName string, filter bson.M, findOptions *moptions.FindOptions, result interface{}) error
	FindOne(ctx context.Context, collectionName string, filter bson.M, result interface{}) error
	Count(ctx context.Context, collectionName string, filter bson.M) (int64, error)
	UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M) error
	DeleteOne(ctx context.Context, collectionName string, filter bson.M) error
}
