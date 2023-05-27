package repository

import (
	"context"
	"golang-starter/service/domain/entity"
)

type SampleRepository interface {
	Create(ctx context.Context, request *entity.SampleMongo) error
	Fetch(ctx context.Context, options map[string]interface{}) ([]entity.SampleMongo, error)
	FetchOne(ctx context.Context, options map[string]interface{}) (entity.SampleMongo, error)
	Count(ctx context.Context, options map[string]interface{}) (int64, error)
	Delete(ctx context.Context, options map[string]interface{}) error
	Update(ctx context.Context, request *entity.SampleMongo) error
}
