package sample

import (
	"context"
	"golang-starter/helpers"
	"golang-starter/service/domain/entity"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func (uc *sampleUC) List(ctx context.Context, options map[string]interface{}) helpers.Response {

	query := options["query"].(url.Values)
	optionRepo := make(map[string]interface{})

	// paginate
	pageQuery := query.Get("page")
	pageInt, err := strconv.Atoi(pageQuery)
	if err != nil {
		pageInt = 1
	} else {
		if pageInt == 0 {
			pageInt = 1
		}
	}
	limitQuery := query.Get("limit")
	limitInt, err := strconv.Atoi(limitQuery)
	if err != nil {
		limitInt = 10
	} else {
		if limitInt == 0 {
			limitInt = 10
		}
	}
	optionRepo["limit"] = int64(limitInt)
	optionRepo["page"] = (int64(pageInt) - 1) * int64(limitInt)

	var model entity.SampleMongo
	// validate allowed_sort
	if query.Get("sort") != "" && helpers.InArrayString(query.Get("sort"), model.AllowedSort()) {
		optionRepo["sort"] = query.Get("sort")
		if query.Get("dir") != "" {
			optionRepo["dir"] = query.Get("dir")
		}
	}

	if id := query.Get("id"); id != "" {
		optionRepo["id"] = id
	}
	sampleMongo, err := uc.SampleRepo.Fetch(ctx, optionRepo)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}

	total, err := uc.SampleRepo.Count(ctx, optionRepo)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err.Error(), err, nil)
	}
	paginate := entity.Paginate{
		List:      sampleMongo,
		Limit:     int64(limitInt),
		Page:      int64(pageInt),
		TotalData: total,
		TotalPage: math.Ceil(float64(total) / float64(limitInt)),
	}
	if total < 1 {
		paginate.List = []string{}
	}
	return helpers.SuccessResponse("success", paginate)
}
