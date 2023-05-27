package usecase

import (
	"context"
	"golang-starter/helpers"
)

type SampleUsecase interface {
	List(ctx context.Context, options map[string]interface{}) helpers.Response
	Create(ctx context.Context, options map[string]interface{}) helpers.Response
	Delete(ctx context.Context, options map[string]interface{}) helpers.Response
	Detail(ctx context.Context, options map[string]interface{}) helpers.Response
	Update(ctx context.Context, options map[string]interface{}) helpers.Response
}
