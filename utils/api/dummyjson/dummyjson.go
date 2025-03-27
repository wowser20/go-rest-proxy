package dummyjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	apiError "go-rest-proxy/internal/errors"
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
		log.Print(err)
		return types.GetDummyProductByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return types.GetDummyProductByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	defer resp.Body.Close()

	var result types.GetDummyProductByIDResponse

	// decode response to json
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
		return types.GetDummyProductByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	return result, nil
}

// GetDummyProducts get all dummy products from dummy json
func GetDummyProducts() (types.GetDummyProductsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products", os.Getenv("DUMMY_JSON_BASE_URL")), nil)
	if err != nil {
		log.Print(err)
		return types.GetDummyProductsResponse{}, errors.New(apiError.DummyJsonError)
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return types.GetDummyProductsResponse{}, errors.New(apiError.DummyJsonError)
	}

	defer resp.Body.Close()

	var result types.GetDummyProductsResponse

	// decode response to json
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
		return types.GetDummyProductsResponse{}, errors.New(apiError.DummyJsonError)
	}

	return result, nil
}

// GetDummyCarts get all dummy carts from dummy json
func GetDummyCarts() (types.GetDummyCartsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/carts", os.Getenv("DUMMY_JSON_BASE_URL")), nil)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartsResponse{}, errors.New(apiError.DummyJsonError)
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartsResponse{}, errors.New(apiError.DummyJsonError)
	}

	defer resp.Body.Close()

	var result types.GetDummyCartsResponse

	// decode response to json
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartsResponse{}, errors.New(apiError.DummyJsonError)
	}

	return result, nil
}

// GetDummyCartByID get dummy cart from dummy json by id
func GetDummyCartByID(cartID int) (types.GetDummyCartByIDResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/carts/%v", os.Getenv("DUMMY_JSON_BASE_URL"), cartID), nil)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	// call dummy json api
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	defer resp.Body.Close()

	var result types.GetDummyCartByIDResponse

	// decode response to json
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
		return types.GetDummyCartByIDResponse{}, errors.New(apiError.DummyJsonError)
	}

	return result, nil
}
