package response

import (
	"encoding/json"
	"golang-starter/service/domain/entity"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultLimit = 10
	defaultPage  = 1
)

type Response struct {
	Status     int                    `json:"status"`
	Message    string                 `json:"message"`
	Validation map[string]interface{} `json:"validation"`
	Data       interface{}            `json:"data"`
}
type Debug struct {
	Property   string
	Error      error
	Additional string
	Function   string
}

func Json(code int, payload interface{}, w http.ResponseWriter) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func ErrorResponse(code int, message string, err error, validation map[string]interface{}) Response {
	return Response{
		Data:       json.NewEncoder(nil),
		Message:    message,
		Status:     code,
		Validation: validation,
	}
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Data:       data,
		Message:    message,
		Status:     http.StatusOK,
		Validation: make(map[string]interface{}),
	}
}

func GeneratePaginateResponse(query url.Values, optionsRepo map[string]interface{}, data interface{}, total int64) entity.Paginate {
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

	paginate := entity.Paginate{
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
