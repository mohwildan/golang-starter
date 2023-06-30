package repository

import (
	"context"
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

func (r *Repository) Find(ctx context.Context, collectionName string, filter bson.M, result interface{}, findOptions *moptions.FindOptions) error {
	collection := r.db.Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindOne(ctx context.Context, collectionName string, filter bson.M, result interface{}) error {
	collection := r.db.Collection(collectionName)
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertOne(ctx context.Context, collectionName string, document interface{}) error {
	if _, err := r.db.Collection(collectionName).InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}
func (r *Repository) UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M) error {
	if _, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteOne(ctx context.Context, collectionName string, filter bson.M) error {
	if _, err := r.db.Collection(collectionName).DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Count(ctx context.Context, collectionName string, filter bson.M) (int64, error) {
	collection := r.db.Collection(collectionName)
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
