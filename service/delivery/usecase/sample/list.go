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
	entity.SetPaginationQuery(query, options, optionsRepo)

	queryArray := []entity.PaginateKey{
		{
			Key:       "id",
			TargetKey: "id",
		},
		{
			Key:       "find_text",
			TargetKey: "text",
		},
	}

	for i := 0; i < len(queryArray); i++ {
		helpers.SetQueryField(query, optionsRepo, queryArray[i])
	}

	var model entity.SampleMongo
	queryOptions, findOptions := helpers.GenerateQuery(optionsRepo)
	err := uc.Repository.Find(ctx, model.GetCollectionName(), queryOptions, &data, findOptions)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	total, err := uc.Repository.Count(ctx, model.GetCollectionName(), queryOptions)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	responsePagination := entity.GeneratePaginateResponse(query, optionsRepo, data, total)

	return helpers.SuccessResponse("success", responsePagination)
}
