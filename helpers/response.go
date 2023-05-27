package helpers

import (
	"encoding/json"
	"net/http"
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
