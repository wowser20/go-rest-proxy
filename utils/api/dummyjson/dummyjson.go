package dummyjson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	client *http.Client = &http.Client{Timeout: 10 * time.Second}
)

// GetDummyProductByID get dummy product from dummy json by id
func GetDummyProductByID(productID string, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products/%s", os.Getenv("DUMMY_JSON_BASE_URL"), productID), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

// GetDummyProducts get all dummy products from dummy json
func GetDummyProducts(response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products", os.Getenv("DUMMY_JSON_BASE_URL")), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
