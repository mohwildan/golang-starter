package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (uc *sampleUC) Update(ctx context.Context, options map[string]interface{}) helpers.Response {
	id, err := helpers.ConvertToObjID(options["id"].(string))
	validation := make(map[string]interface{})
	request := options["request"].(entity.SampleMongo)
	filter := map[string]interface{}{
		"id": id,
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrorResponse(http.StatusNotFound, "sample tidak ditemukan", err, nil)
		}
	}
	var row entity.SampleMongo
	err = uc.Repository.FindOne(ctx, row.GetCollectionName(), filter, &row)

	now := time.Now().UTC()
	row.UpdatedAt = &now
	if row.Text != "" {
		row.Text = request.Text
	}

	update := bson.M{
		"$set": &row,
	}

	err = uc.Repository.UpdateOne(ctx, row.GetCollectionName(), filter, update)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, validation)
	}
	return helpers.SuccessResponse("success", map[string]interface{}{})
}
