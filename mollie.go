package mollie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	Secret  string
	Version string
}

// makes http request to the api using api key secret as authorization
func (c *APIClient) request(endpoint string, method string, data []byte) ([]byte, error) {
	client := &http.Client{}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest(method, "https://api.mollie.com/"+c.Version+"/"+endpoint, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Secret)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// API request succeeded
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return raw, nil
	}
	err = extractError(res.StatusCode, raw)
	return nil, err
}

type Link struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type Error struct {
	Status int             `json:"status"`
	Title  string          `json:"title"`
	Detail string          `json:"detail"`
	Links  map[string]Link `json:"_links"`
}

// extract error data from mollie response
func extractError(status int, raw []byte) error {
	if status >= 400 && status < 500 {
		e := Error{}
		if err := json.Unmarshal(raw, &e); err != nil {
			return fmt.Errorf("failed to unmarshal json: %v", err)
		}
		return fmt.Errorf("[error code %v] %v: %v", e.Status, e.Title, e.Detail)
	}
	if status >= 500 && status < 600 {
		e := Error{}
		if err := json.Unmarshal(raw, &e); err != nil {
			return fmt.Errorf("failed to unmarshal json: %v", err)
		}
		return fmt.Errorf("[error code %v] %v: %v (THIS IS A MOLLIE SERVER ERROR)", e.Status, e.Title, e.Detail)
	}
	// just return the full JSON body and the status int
	return fmt.Errorf("unknown error status %v, return body: %v", status, string(raw))
}

// safely get data at pointer
func unpoint[T any](t *T) T {
	if t == nil {
		var result T
		return result
	}
	return *t
}

type ListResponse struct {
	Count    int                        `json:"count"`
	Embedded map[string]json.RawMessage `json:"_embedded"`
}

// Finish later
func toQueryString(param any) {
	raw, _ := json.Marshal(param)
	paramsMAP := map[string]string{}
	json.Unmarshal(raw, &paramsMAP)
	query := ""
	for key, val := range paramsMAP {
		query = query + fmt.Sprintf("&%v=%v", key, val)
	}
}
