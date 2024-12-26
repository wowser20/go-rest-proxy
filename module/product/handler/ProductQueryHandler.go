package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/utils/api/dummyjson"
)

// GetDummyProducts gets all products
func GetDummyProducts(w http.ResponseWriter, r *http.Request) {
	result, err := dummyjson.GetDummyProducts()
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Error loading products.",
			ErrorCode: errors.DummyJsonError,
		}

		response.JSON(w)
		return
	}

	response := &viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched all products.",
		Data:    result,
	}

	response.JSON(w)
}

// GetDummyProductByID gets a product by id
func GetDummyProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid product ID.",
			ErrorCode: errors.InvalidPayload,
		}

		response.JSON(w)
		return
	}

	result, err := dummyjson.GetDummyProductByID(productID)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Error loading product by id.",
			ErrorCode: errors.DummyJsonError,
		}

		response.JSON(w)
		return
	}

	response := &viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched product by id.",
		Data:    result,
	}

	response.JSON(w)
}
