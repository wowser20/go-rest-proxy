package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/module/product/handler/types"
	"go-rest-proxy/utils/api/dummyjson"
)

// GetDummyProductsHandler handler for getting all products from dummyjson
func GetDummyProductsHandler(w http.ResponseWriter, r *http.Request) {
	var result types.GetProductsResponse

	err := dummyjson.GetDummyProducts(&result)
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

// GetDummyProductsHandler handler for getting all products from dummyjson
func GetDummyProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productID")

	var result types.GetProductByIDResponse

	err := dummyjson.GetDummyProductByID(id, &result)
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
