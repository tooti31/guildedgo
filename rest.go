package guildedgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (c *Client) PostRequest(endpoint string, body interface{}) ([]byte, error) {
	jsonBody, _ := json.Marshal(&body)

	resp, err := DoRequest("POST", endpoint, jsonBody, c.Token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) PostRequestV2(endpoint string, body any, v any) error {
	jsonBody, _ := json.Marshal(&body)

	resp, err := DoRequest("POST", endpoint, jsonBody, c.Token)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PatchRequest(endpoint string, body any, v any) error {
	jsonBody, _ := json.Marshal(&body)

	resp, err := DoRequest("PATCH", endpoint, jsonBody, c.Token)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRequest(endpoint string) ([]byte, error) {
	resp, err := DoRequest("GET", endpoint, nil, c.Token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetRequestV2(endpoint string, v any) error {
	resp, err := DoRequest("GET", endpoint, nil, c.Token)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PutRequest(endpoint string, body interface{}) ([]byte, error) {
	jsonBody, _ := json.Marshal(&body)

	resp, err := DoRequest("PUT", endpoint, jsonBody, c.Token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) PutRequestV2(endpoint string, body interface{}, v any) error {
	jsonBody, _ := json.Marshal(&body)

	resp, err := DoRequest("PUT", endpoint, jsonBody, c.Token)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteRequest(endpoint string) ([]byte, error) {
	resp, err := DoRequest("DELETE", endpoint, nil, c.Token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

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
	case http.StatusBadGateway, http.StatusForbidden, http.StatusBadRequest, http.StatusNotFound:
		var resError responseError

		err := json.Unmarshal(body, &resError)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(resError.Message)
	}

	return body, nil
}
