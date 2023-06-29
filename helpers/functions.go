package helpers

import (
	"net/url"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func SetQueryField(query url.Values, optionRepo map[string]interface{}, fieldName string) {
	if value := query.Get(fieldName); value != "" {
		optionRepo[fieldName] = value
	}
}

const (
	defaultLimit   = 10
	defaultPage    = 1
	defaultSortBy  = "created_at"
	defaultSortDir = "asc"
)

func GenerateQuery(options map[string]interface{}) (bson.M, *moptions.FindOptions) {
	query := bson.M{}

	if id, ok := options["id"].(primitive.ObjectID); ok {
		query["_id"] = id
	} else if id, ok := options["id"].(string); ok {
		objID, _ := primitive.ObjectIDFromHex(id)
		query["_id"] = objID
	}

	page := getOptionAsInt64(options, "page", defaultPage)
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
func ConvertToObjID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}
