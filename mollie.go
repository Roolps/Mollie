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

	err = extractError(raw)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

type Link struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type Error struct {
	Status *int    `json:"status"`
	Title  *string `json:"title"`
	Detail *string `json:"detail"`
	Links  *[]Link `json:"_links"`
}

// extract error data from mollie response
func extractError(raw []byte) error {
	err := Error{}
	json.Unmarshal(raw, &err)
	if err.Status != nil && err.Title != nil && err.Detail != nil && err.Links != nil {
		return fmt.Errorf("[error code %v] %v: %v", unpoint(err.Status), unpoint(err.Title), unpoint(err.Detail))
	}
	return nil
}

// safely get data at pointer
func unpoint[T any](t *T) T {
	if t == nil {
		var result T
		return result
	}
	return *t
}
