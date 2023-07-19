package helpers

import (
	"golang-starter/service/domain/entity"
	"net/url"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func SetQueryField(query url.Values, optionRepo map[string]interface{}, fieldKey entity.PaginateKey) {
	if value := query.Get(fieldKey.QueryKey); value != "" {
		optionRepo[fieldKey.TargetQueryKey] = value
	}
}

func GenerateQuery(options map[string]interface{}) (bson.M, *moptions.FindOptions) {
	const (
		defaultLimit   = 10
		defaultPage    = 1
		defaultSort    = "asc"
		defaultSortDir = "created_at"
	)
	query := bson.M{}

	mongoOptions := moptions.Find()
	for key, value := range options {
		switch key {
		case "id":
			if id, ok := value.(primitive.ObjectID); ok {
				query["_id"] = id
			} else if id, ok := value.(string); ok {
				objID, _ := ConvertToObjID(id)
				query["_id"] = objID
			}
		case "limit", "page", "sort", "dir":
			page := getOptionAsInt64(options, "page", defaultPage)
			limit := getOptionAsInt64(options, "limit", defaultLimit)
			sort := getOptionAsString(options, "sort", defaultSort)
			sortBy := getOptionAsString(options, "dir", defaultSortDir)

			mongoOptions.SetSkip(page * limit)
			mongoOptions.SetLimit(limit)

			sortQ := bson.M{}
			if strings.ToLower(sort) == "desc" {
				sortQ[sortBy] = -1
			} else {
				sortQ[sortBy] = 1
			}
			mongoOptions.SetSort(sortQ)
		default:
			if _, exists := query[key]; !exists {
				query[key] = value
			}
		}
	}

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
