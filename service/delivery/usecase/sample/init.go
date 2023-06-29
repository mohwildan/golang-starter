package sample

import (
	"golang-starter/service/domain/repository"
	"golang-starter/service/domain/usecase"
	"time"
)

type sampleUC struct {
	contextTimeout time.Duration
	Repository     repository.Repository
}

func Usecase(contextTimeout time.Duration, repository repository.Repository) usecase.SampleUsecase {
	return &sampleUC{
		contextTimeout: contextTimeout,
		Repository:     repository,
	}
}
