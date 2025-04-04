package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/utils/api/dummyjson"
)

// GetCarts gets all carts
func GetCarts(w http.ResponseWriter, r *http.Request) {
	res, err := dummyjson.GetDummyCarts()
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DummyJsonError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error loading carts."
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
		Message: "Successfully fetched all carts.",
		Data:    res,
	}

	response.JSON(w)
}

// GetCartByID gets a cart by id
func GetCartByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cartID, err := strconv.Atoi(id)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Invalid cart ID.",
			ErrorCode: errors.InvalidPayload,
		}

		response.JSON(w)
		return
	}

	res, err := dummyjson.GetDummyCartByID(cartID)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DummyJsonError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error loading cart by id."
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
		Message: "Successfully fetched cart by id.",
		Data:    res,
	}

	response.JSON(w)
}
