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
	id := options["id"].(string)

	var sample entity.SampleMongo
	validation := make(map[string]interface{})
	filter, _ := helpers.GenerateQuery(bson.M{
		"id": id,
	})
	err := uc.Repository.FindOne(ctx, sample.GetCollectionName(), filter, &sample)
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
