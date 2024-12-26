package dummyjson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go-rest-proxy/utils/api/dummyjson/types"
)

// client timeout
var (
	client *http.Client = &http.Client{Timeout: 10 * time.Second}
)

// GetDummyProductByID get dummy product from dummy json by id
func GetDummyProductByID(productID int) (types.GetDummyProductByIDResponse, error) {
	// call dummy json api
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products/%v", os.Getenv("DUMMY_JSON_BASE_URL"), productID), nil)
	if err != nil {
		return types.GetDummyProductByIDResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return types.GetDummyProductByIDResponse{}, err
	}

	defer resp.Body.Close()

	var result types.GetDummyProductByIDResponse

	// decode response to json
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.GetDummyProductByIDResponse{}, err
	}

	return result, nil
}

// GetDummyProducts get all dummy products from dummy json
func GetDummyProducts() (types.GetDummyProductsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products", os.Getenv("DUMMY_JSON_BASE_URL")), nil)
	if err != nil {
		return types.GetDummyProductsResponse{}, err
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		return types.GetDummyProductsResponse{}, err
	}

	defer resp.Body.Close()

	var result types.GetDummyProductsResponse

	// decode response to json
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.GetDummyProductsResponse{}, err
	}

	return result, nil
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
