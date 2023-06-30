package entity

import (
	"math"
	"net/url"
	"strconv"
)

const (
	defaultLimit = 10
	defaultPage  = 1
)

type Paginate struct {
	List      interface{} `json:"list"`
	Limit     int64       `json:"limit"`
	Page      int64       `json:"page"`
	TotalData int64       `json:"total_data"`
	TotalPage float64     `json:"total_page"`
}
type PaginateKey struct {
	Key       string
	TargetKey string
}

func SetPaginationQuery(query url.Values, options map[string]interface{}, optionsRepo map[string]interface{}) {
	//paginate
	pageQuery := query.Get("page")
	pageInt, err := strconv.Atoi(pageQuery)
	if err != nil || pageInt <= 0 {
		pageInt = defaultPage
	}

	limitQuery := query.Get("limit")
	limitInt, err := strconv.Atoi(limitQuery)
	if err != nil || limitInt <= 0 {
		limitInt = defaultLimit
	}
	optionsRepo["limit"] = int64(limitInt)
	optionsRepo["page"] = (int64(pageInt) - 1) * int64(limitInt)

	// Validate allowed_sort
	if sort := query.Get("sort"); sort != "" {
		optionsRepo["sort"] = sort
		if dir := query.Get("dir"); dir != "" {
			optionsRepo["dir"] = dir
		}
	}
}

func GeneratePaginateResponse(query url.Values, optionsRepo map[string]interface{}, data interface{}, total int64) Paginate {
	// Paginate
	pageQuery := query.Get("page")
	pageInt, err := strconv.Atoi(pageQuery)
	if err != nil || pageInt <= 0 {
		pageInt = defaultPage
	}

	limitQuery := query.Get("limit")
	limitInt, err := strconv.Atoi(limitQuery)
	if err != nil || limitInt <= 0 {
		limitInt = defaultLimit
	}

	paginate := Paginate{
		List:      data,
		Limit:     int64(limitInt),
		Page:      int64(pageInt),
		TotalData: total,
		TotalPage: math.Ceil(float64(total) / float64(limitInt)),
	}

	if total < 1 {
		paginate.List = []string{}
	}

	return paginate
}
