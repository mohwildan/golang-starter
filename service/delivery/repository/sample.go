package repository

import (
	"context"
	"errors"
	"golang-starter/service/domain/entity"
	"golang-starter/service/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var ErrSampleNotFound = errors.New("Sample not found")

type sampleRepo struct {
	db *mongo.Database
}

func SampleRepo(db *mongo.Database) repository.SampleRepository {
	return &sampleRepo{
		db: db,
	}
}

func (r *sampleRepo) Create(ctx context.Context, sample *entity.SampleMongo) error {
	if _, err := r.db.Collection(sample.GetCollectionName()).InsertOne(ctx, *sample); err != nil {
		return errors.New("failed to create sample: " + err.Error())
	}
	return nil
}

func (r *sampleRepo) FetchOne(ctx context.Context, options map[string]interface{}) (entity.SampleMongo, error) {
	sample := entity.SampleMongo{}
	query, _ := sample.GenerateQuery(options)
	if err := r.db.Collection(sample.GetCollectionName()).FindOne(ctx, query).Decode(&sample); err != nil {
		if err == mongo.ErrNoDocuments {
			return sample, ErrSampleNotFound
		}
		return sample, errors.New("failed to fetch sample: " + err.Error())
	}
	return sample, nil
}

func (r *sampleRepo) Fetch(ctx context.Context, options map[string]interface{}) ([]entity.SampleMongo, error) {
	model := entity.SampleMongo{}
	query, findOptions := model.GenerateQuery(options)
	cursor, err := r.db.Collection(model.GetCollectionName()).Find(ctx, query, findOptions)
	if err != nil {
		return nil, errors.New("failed to fetch samples: " + err.Error())
	}
	defer cursor.Close(ctx)

	samples := make([]entity.SampleMongo, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		sample := entity.SampleMongo{}
		if err := cursor.Decode(&sample); err != nil {
			return samples, errors.New("failed to decode sample: " + err.Error())
		}
		samples = append(samples, sample)
	}
	if err := cursor.Err(); err != nil {
		return samples, errors.New("cursor error: " + err.Error())
	}
	return samples, nil
}

func (r *sampleRepo) Count(ctx context.Context, options map[string]interface{}) (int64, error) {
	model := entity.SampleMongo{}
	query, _ := model.GenerateQuery(options)
	total, err := r.db.Collection(model.GetCollectionName()).CountDocuments(ctx, query)
	if err != nil {
		return 0, errors.New("failed to count samples: " + err.Error())
	}
	return total, nil
}

func (r *sampleRepo) Delete(ctx context.Context, options map[string]interface{}) error {
	model := entity.SampleMongo{}
	query, _ := model.GenerateQuery(options)
	result, err := r.db.Collection(model.GetCollectionName()).DeleteOne(ctx, query)
	if err != nil {
		return errors.New("failed to delete sample: " + err.Error())
	}
	if result.DeletedCount == 0 {
		return ErrSampleNotFound
	}
	return nil
}

func (r *sampleRepo) Update(ctx context.Context, sample *entity.SampleMongo) error {
	filter := bson.M{"_id": sample.ID}
	update := bson.M{"$set": sample}
	result, err := r.db.Collection(sample.GetCollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to update sample: " + err.Error())
	}
	if result.ModifiedCount == 0 {
		return ErrSampleNotFound
	}
	return nil
}
