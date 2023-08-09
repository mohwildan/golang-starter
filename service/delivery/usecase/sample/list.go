package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"net/http"
	"net/url"
)

func (uc *sampleUC) List(ctx context.Context, options map[string]interface{}) helpers.Response {
	query := options["query"].(url.Values)
	optionsRepo := make(map[string]interface{})
	var data []entity.SampleMongo

	// Paginate Query
	helpers.SetPaginationQuery(query, optionsRepo)

	var filters []helpers.Filter
	if v := query.Get("text"); v != "" {
		filters = append(filters, helpers.Filter{Field: "text", Operator: helpers.Contains, Value: v})
	}
	if v := query.Get("id"); v != "" {
		objId, _ := helpers.ConvertToObjID(v)
		filters = append(filters, helpers.Filter{Field: "_id", Operator: helpers.Equal, Value: objId})
	}

	processQuery := helpers.GenerateQuery(filters, optionsRepo)
	err := uc.Repository.Find(ctx, helpers.SampleCollectionName, processQuery.Query, &data, processQuery.FindOptions)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	total, err := uc.Repository.Count(ctx, helpers.SampleCollectionName, filters)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	responsePagination := helpers.GeneratePaginateResponse(query, optionsRepo, data, total)

	return helpers.SuccessResponse("success", responsePagination)
}
