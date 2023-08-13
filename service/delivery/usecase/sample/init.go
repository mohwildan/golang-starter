package sample

import (
	"context"
	"golang-starter/helpers/response"
	"golang-starter/service/domain/repository"
	"golang-starter/service/domain/usecase"
	"time"
)

type sampleUC struct {
	contextTimeout time.Duration
	repository     repository.Repository
}

func (uc sampleUC) Create(ctx context.Context, options map[string]interface{}) response.Response {
	//TODO implement me
	panic("implement me")
}

func (uc sampleUC) Delete(ctx context.Context, options map[string]interface{}) response.Response {
	//TODO implement me
	panic("implement me")
}

func (uc sampleUC) Detail(ctx context.Context, options map[string]interface{}) response.Response {
	//TODO implement me
	panic("implement me")
}

func (uc sampleUC) Update(ctx context.Context, options map[string]interface{}) response.Response {
	//TODO implement me
	panic("implement me")
}

func NewUsecase(contextTimeout time.Duration, repository repository.Repository) usecase.SampleUsecase {
	return &sampleUC{
		contextTimeout: contextTimeout,
		repository:     repository,
	}
}
