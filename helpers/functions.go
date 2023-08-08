package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOptions[T any](options map[string]any, key string, defaultValue T) T {
	if value, ok := options[key].(T); ok {
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
