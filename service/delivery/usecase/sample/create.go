package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *sampleUC) Create(ctx context.Context, options map[string]interface{}) helpers.Response {
	request := options["request"].(entity.SampleMongo)

	now := time.Now().UTC()
	sample := entity.SampleMongo{
		ID:        primitive.NewObjectID(),
		Text:      request.Text,
		CreatedAt: now,
		UpdatedAt: nil,
	}

	err := uc.SampleRepo.Create(ctx, &sample)

	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	return helpers.SuccessResponse("success", map[string]interface{}{})
}
