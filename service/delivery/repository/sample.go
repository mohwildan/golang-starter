package repository

import (
	"context"
	"golang-starter/service/domain/entity"
	"golang-starter/service/domain/repository"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type sampleRepo struct {
	db *mongo.Database
}

func SampleRepo(db *mongo.Database) repository.SampleRepository {
	return &sampleRepo{
		db: db,
	}
}

func (r *sampleRepo) Create(ctx context.Context, sample *entity.SampleMongo) error {
	_, err := r.db.Collection(sample.GetCollectionName()).InsertOne(context.TODO(), *sample)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (r *sampleRepo) FetchOne(ctx context.Context, options map[string]interface{}) (sample entity.SampleMongo, err error) {
	query, _ := sample.GenerateQuery(options)
	err = r.db.Collection(sample.GetCollectionName()).FindOne(ctx, query).Decode(&sample)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (r *sampleRepo) Fetch(ctx context.Context, options map[string]interface{}) (list []entity.SampleMongo, err error) {
	model := entity.SampleMongo{}
	query, findOptions := model.GenerateQuery(options)
	cursor, err := r.db.Collection(model.GetCollectionName()).Find(ctx, query, findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		row := entity.SampleMongo{}
		err = cursor.Decode(&row)
		if err != nil {
			log.Println(err.Error())
			return list, err
		}
		list = append(list, row)
	}
	return list, nil
}
func (r *sampleRepo) Count(ctx context.Context, options map[string]interface{}) (int64, error) {
	var model entity.SampleMongo

	query, _ := model.GenerateQuery(options)
	total, err := r.db.Collection(model.GetCollectionName()).CountDocuments(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}
	return total, nil

}
func (r *sampleRepo) Delete(ctx context.Context, options map[string]interface{}) error {
	var model entity.SampleMongo
	query, _ := model.GenerateQuery(options)
	_, err := r.db.Collection(model.GetCollectionName()).DeleteOne(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func (r *sampleRepo) Update(ctx context.Context, sample *entity.SampleMongo) error {
	_, err := r.db.Collection(sample.GetCollectionName()).UpdateOne(ctx, bson.M{
		"_id": sample.ID,
	},
		bson.M{
			"$set": &sample,
		})
	if err != nil {
		return err
	}
	return nil
}
