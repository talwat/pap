package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func Get(url string, bodyOut interface{}) int {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	Error(err, "an error occurred while making request")

	resp, err := http.DefaultClient.Do(req)
	Error(err, "an error occurred while sending request to papermc api")

	err = json.NewDecoder(resp.Body).Decode(&bodyOut)
	Error(err, "an error occurred while decoding response")

	defer resp.Body.Close()

	return resp.StatusCode
}
