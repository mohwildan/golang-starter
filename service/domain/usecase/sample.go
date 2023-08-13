package usecase

import (
	"context"
	"golang-starter/helpers/response"
)

type SampleUsecase interface {
	List(ctx context.Context, options map[string]interface{}) response.Response
	Create(ctx context.Context, options map[string]interface{}) response.Response
	Delete(ctx context.Context, options map[string]interface{}) response.Response
	Detail(ctx context.Context, options map[string]interface{}) response.Response
	Update(ctx context.Context, options map[string]interface{}) response.Response
}
