package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang-starter/helpers"

	moptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertOne(ctx context.Context, collectionName string, document any) error
	Find(ctx context.Context, collectionName string, filter bson.M, result any, findOptions *moptions.FindOptions) error
	FindOne(ctx context.Context, collectionName string, filters []helpers.Filter, result any) error
	Count(ctx context.Context, collectionName string, filters []helpers.Filter) (int64, error)
	UpdateOne(ctx context.Context, collectionName string, filters []helpers.Filter, update bson.M) error
	DeleteOne(ctx context.Context, collectionName string, filters []helpers.Filter) error
}
