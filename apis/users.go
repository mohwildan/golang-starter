package apis

import (
	"encoding/json"
	"fmt"
)

type UserApiService struct {
	api *APIService
}

func NewUserApiService(baseURL string, customHeaders map[string]string) *UserApiService {
	api := NewAPIService(baseURL, customHeaders)
	return &UserApiService{
		api: api,
	}
}

type UserDetail struct {
	Name string `json:"name"`
}

func (u *UserApiService) GetDetail(id string) (*UserDetail, error) {
	endpoint := fmt.Sprintf("/v2/pokemon/%s", id)
	response, err := u.api.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var userDetail UserDetail
	if err := json.Unmarshal(response, &userDetail); err != nil {
		return nil, err
	}

	return &userDetail, nil
}

type LogParams struct {
	Data struct {
		Attributes  map[string]interface{} `json:"attributes,omitempty"`
		Description string                 `json:"description,omitempty"`
		Type        string                 `json:"type,omitempty"`
		SubType     string                 `json:"sub_type,omitempty"`
	} `json:"data,omitempty"`
}

type LogResponse struct {
}

func (u *UserApiService) CreateLog(params LogParams) (*LogResponse, error) {
	endpoint := "/user/v1/log/create"
	response, err := u.api.SendRequest("POST", endpoint, params)
	if err != nil {
		return nil, err
	}

	var logResponse LogResponse
	if err := json.Unmarshal(response, &logResponse); err != nil {
		return nil, err
	}

	return &logResponse, nil
}
