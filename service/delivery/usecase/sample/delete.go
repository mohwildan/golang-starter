package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (uc *sampleUC) Delete(ctx context.Context, options map[string]interface{}) helpers.Response {
	id, err := helpers.ConvertToObjID(options["id"].(string))

	if err != nil {
		return helpers.ErrorResponse(http.StatusNotFound, err.Error(), err, nil)
	}

	var sample entity.SampleMongo
	validation := make(map[string]interface{})
	filter := bson.M{
		"_id": id,
	}
	err = uc.Repository.FindOne(ctx, sample.GetCollectionName(), filter, &sample)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "sample tidak di timukan", err, validation)
		}
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, validation)
	}
	err = uc.Repository.DeleteOne(ctx, sample.GetCollectionName(), filter)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}
	return helpers.SuccessResponse("success", map[string]interface{}{})
}
