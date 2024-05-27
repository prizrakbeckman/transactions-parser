package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchDataFromUrl(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non 200 response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
