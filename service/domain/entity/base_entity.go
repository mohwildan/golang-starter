package entity

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type BaseEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
}

const (
	defaultLimit   = 10
	defaultSortBy  = "created_at"
	defaultSortDir = "asc"
)

func (e *BaseEntity) GenerateQuery(options map[string]interface{}) (bson.M, *moptions.FindOptions) {
	query := bson.M{}

	if id, ok := options["id"].(primitive.ObjectID); ok {
		query["_id"] = id
	} else if id, ok := options["id"].(string); ok {
		objID, _ := primitive.ObjectIDFromHex(id)
		query["_id"] = objID
	}

	page := getOptionAsInt64(options, "page", 0)
	limit := getOptionAsInt64(options, "limit", defaultLimit)
	sortBy := getOptionAsString(options, "sort", defaultSortBy)
	sortDir := getOptionAsString(options, "dir", defaultSortDir)

	mongoOptions := moptions.Find()
	mongoOptions.SetSkip(page * limit)
	mongoOptions.SetLimit(limit)

	sortQ := bson.M{}
	if strings.ToLower(sortDir) == "desc" {
		sortQ[sortBy] = -1
	} else {
		sortQ[sortBy] = 1
	}
	mongoOptions.SetSort(sortQ)

	return query, mongoOptions
}

func (e *BaseEntity) AllowedSort() []string {
	// Menyesuaikan dengan bidang yang dapat diurutkan pada entitas yang bersangkutan
	return []string{"_id", "created_at", "updated_at"}
}

func getOptionAsInt64(options map[string]interface{}, key string, defaultValue int64) int64 {
	if value, ok := options[key].(int64); ok {
		return value
	}
	return defaultValue
}

func getOptionAsString(options map[string]interface{}, key string, defaultValue string) string {
	if value, ok := options[key].(string); ok {
		return value
	}
	return defaultValue
}
