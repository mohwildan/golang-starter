package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SampleMongo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
}

type RequestSampleUpdate struct {
	Text string `json:"text,omitempty"`
}

func (s *SampleMongo) GetCollectionName() string {
	return "sample"
}
