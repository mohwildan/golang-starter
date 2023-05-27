package sample

import (
	"context"
	"golang-starter/helpers"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (uc *sampleUC) Delete(ctx context.Context, options map[string]interface{}) helpers.Response {
	id := options["id"].(string)

	optionRepo := map[string]interface{}{
		"id": id,
	}
	validation := make(map[string]interface{})
	_, err := uc.SampleRepo.FetchOne(ctx, optionRepo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "sample tidak di timukan", err, validation)
		}
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, validation)
	}
	err = uc.SampleRepo.Delete(ctx, optionRepo)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}
	return helpers.SuccessResponse("success", map[string]interface{}{})
}
