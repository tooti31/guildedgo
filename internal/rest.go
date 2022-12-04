package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type responseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func DoRequest(method string, endpoint string, body []byte, token string) ([]byte, error) {
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	// Handle status codes - remember rate limiting
	resp, err := do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusBadGateway:
	case http.StatusForbidden:
	case http.StatusBadRequest:
		var resError responseError

		err := json.Unmarshal(body, &resError)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(resError.Message)
	}

	return body, nil
}
