package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (uc *sampleUC) Update(ctx context.Context, options map[string]interface{}) helpers.Response {
	id := options["id"].(string)
	validation := make(map[string]interface{})
	request := options["request"].(entity.SampleMongo)
	objId, err := primitive.ObjectIDFromHex(id)
	optionRepo := map[string]interface{}{
		"id": objId,
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "sample tidak ditemukan", err, nil)
		}
	}
	row, _ := uc.SampleRepo.FetchOne(ctx, optionRepo)

	now := time.Now().UTC()
	row.UpdatedAt = &now
	if row.Text != "" {
		row.Text = request.Text
	}

	err = uc.SampleRepo.Update(ctx, &row)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, validation)
	}
	return helpers.SuccessResponse("success", map[string]interface{}{})
}
