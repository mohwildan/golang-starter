package sample

import (
	"golang-starter/service/domain/repository"
	"golang-starter/service/domain/usecase"
	"time"
)

type sampleUC struct {
	contextTimeout time.Duration
	SampleRepo     repository.SampleRepository
}

func Usecase(sampleRepo repository.SampleRepository, contextTimeout time.Duration) usecase.SampleUsecase {
	return &sampleUC{
		contextTimeout: contextTimeout,
		SampleRepo:     sampleRepo,
	}
}
