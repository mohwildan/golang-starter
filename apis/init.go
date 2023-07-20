package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type APIService struct {
	BaseURL       string
	CustomHeaders map[string]string
}

func NewAPIService(baseURL string, customHeaders map[string]string) *APIService {
	return &APIService{
		BaseURL:       baseURL,
		CustomHeaders: customHeaders,
	}
}

func (api *APIService) SendRequest(method, endpoint string, body interface{}) ([]byte, error) {
	url := api.BaseURL + endpoint

	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	for key, value := range api.CustomHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Request failed with status: %s, response: %s", resp.Status, string(responseBody))
	}

	return responseBody, nil
}
