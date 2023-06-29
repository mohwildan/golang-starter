package entity

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type SampleMongo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt time.Time          `json"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json"updated_at" bson:"updated_at"`
}

func (s *SampleMongo) GetCollectionName() string {
	return "sample"
}

func (s *SampleMongo) GenerateQuery(options map[string]interface{}) (bson.M, *moptions.FindOptions) {
	query := bson.M{}

	if id, ok := options["id"].(primitive.ObjectID); ok {
		query["_id"] = id
	} else if id, ok := options["id"].(string); ok {
		objId, _ := primitive.ObjectIDFromHex(id)
		query["_id"] = objId

	}

	// limit, page & sort
	page, ok := options["page"].(int64)
	if !ok {
		page = 0
	}
	limit, ok := options["limit"].(int64)
	if !ok {
		limit = 10
	}

	sortBy, ok := options["sort"].(string)
	if !ok {
		sortBy = "created_at"
	}
	sortDir, ok := options["dir"].(string)
	if !ok {
		sortDir = "asc"
	}
	mongoOptions := moptions.Find()
	mongoOptions.SetSkip(page)
	mongoOptions.SetLimit(limit)
	if sortBy != "" {
		sortQ := bson.M{}
		sortDirMongo := int(1)
		if strings.ToLower(sortDir) == "desc" {
			sortDirMongo = -1
		}
		sortQ[sortBy] = sortDirMongo
		mongoOptions.SetSort(sortQ)
	}

	return query, mongoOptions
}

func (s *SampleMongo) AllowedSort() []string {
	return []string{"_id", "status", "created_at", "updated_at"}
}
