package entity

type SampleMongo struct {
	BaseEntity
	Text string `json:"text" bson:"text"`
}

func (s *SampleMongo) GetCollectionName() string {
	return "sample"
}
