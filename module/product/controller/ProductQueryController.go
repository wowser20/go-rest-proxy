package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/module/product/service"
)

type ProductQueryController struct {
	service.ProductQueryService
}

// GetProducts gets all products
func (controller *ProductQueryController) GetProducts(w http.ResponseWriter, r *http.Request) {
	result, err := controller.ProductQueryService.GetProducts()
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DummyJsonError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error loading products."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
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

// GetProductByID gets a product by id
func (controller *ProductQueryController) GetProductByID(w http.ResponseWriter, r *http.Request) {
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

	result, err := controller.ProductQueryService.GetProductByID(productID)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DummyJsonError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error loading product by id."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
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
