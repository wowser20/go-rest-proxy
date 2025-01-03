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

// GetDummyCarts get all dummy carts from dummy json
func GetDummyCarts() (types.GetDummyCartsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/carts", os.Getenv("DUMMY_JSON_BASE_URL")), nil)
	if err != nil {
		return types.GetDummyCartsResponse{}, err
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		return types.GetDummyCartsResponse{}, err
	}

	defer resp.Body.Close()

	var result types.GetDummyCartsResponse

	// decode response to json
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.GetDummyCartsResponse{}, err
	}

	return result, nil
}

// GetDummyCartByID get dummy cart from dummy json by id
func GetDummyCartByID(cartID int) (types.GetDummyCartByIDResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/carts/%v", os.Getenv("DUMMY_JSON_BASE_URL"), cartID), nil)
	if err != nil {
		return types.GetDummyCartByIDResponse{}, err
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		return types.GetDummyCartByIDResponse{}, err
	}

	defer resp.Body.Close()

	var result types.GetDummyCartByIDResponse

	// decode response to json
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.GetDummyCartByIDResponse{}, err
	}

	return result, nil
}
