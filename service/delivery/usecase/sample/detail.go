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
	id := options["id"].(string)
	objId, _ := helpers.ConvertToObjID(id)
	filters := []helpers.Filter{
		{"_id", helpers.Equal, objId},
	}
	err := uc.Repository.FindOne(ctx, helpers.SampleCollectionName, filters, &sample)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "data sample tidak di temukan", err, nil)
		}
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	return helpers.SuccessResponse("success", sample)
}
