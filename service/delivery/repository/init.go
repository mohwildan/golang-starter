package repository

import (
	"context"
	"fmt"
	"golang-starter/helpers"
	"golang-starter/service/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) repository.Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context, collectionName string, filter bson.M, result any, findOptions *moptions.FindOptions) error {
	collection := r.db.Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}(cursor, ctx)

	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindOne(ctx context.Context, collectionName string, filters []helpers.Filter, result any) error {
	collection := r.db.Collection(collectionName)
	processQuery := helpers.GenerateQuery(filters, bson.M{})
	err := collection.FindOne(ctx, processQuery.Query).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertOne(ctx context.Context, collectionName string, document any) error {
	if _, err := r.db.Collection(collectionName).InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}
func (r *Repository) UpdateOne(ctx context.Context, collectionName string, filters []helpers.Filter, update bson.M) error {
	processQuery := helpers.GenerateQuery(filters, bson.M{})
	if _, err := r.db.Collection(collectionName).UpdateOne(ctx, processQuery.Query, update); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteOne(ctx context.Context, collectionName string, filters []helpers.Filter) error {
	processQuery := helpers.GenerateQuery(filters, bson.M{})
	if _, err := r.db.Collection(collectionName).DeleteOne(ctx, processQuery.Query); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Count(ctx context.Context, collectionName string, filters []helpers.Filter) (int64, error) {
	processQuery := helpers.GenerateQuery(filters, bson.M{})
	collection := r.db.Collection(collectionName)
	count, err := collection.CountDocuments(ctx, processQuery.Query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
