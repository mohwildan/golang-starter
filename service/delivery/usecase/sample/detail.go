package sample

import (
	"context"
	"golang-starter/helpers"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (uc *sampleUC) Detail(ctx context.Context, options map[string]interface{}) helpers.Response {
	id := options["id"].(string)
	optionRepo := map[string]interface{}{
		"id": id,
	}

	sample, err := uc.SampleRepo.FetchOne(ctx, optionRepo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "data sample tidak di temukan", err, nil)
		}
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	return helpers.SuccessResponse("success", sample)
}
