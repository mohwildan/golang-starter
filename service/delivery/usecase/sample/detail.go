package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (uc *sampleUC) Detail(ctx context.Context, options map[string]interface{}) helpers.Response {
	var sample entity.SampleMongo
	id, err := helpers.ConvertToObjID(options["id"].(string))
	filter := map[string]interface{}{
		"_id": id,
	}
	err = uc.Repository.FindOne(ctx, sample.GetCollectionName(), filter, &sample)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "data sample tidak di temukan", err, nil)
		}
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	return helpers.SuccessResponse("success", sample)
}
