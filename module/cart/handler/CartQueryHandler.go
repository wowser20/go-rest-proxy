package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/utils/api/dummyjson"
)

// GetDummyCarts gets all carts
func GetDummyCarts(w http.ResponseWriter, r *http.Request) {
	result, err := dummyjson.GetDummyCarts()
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Error loading carts.",
			ErrorCode: errors.DummyJsonError,
		}

		response.JSON(w)
		return
	}

	response := &viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched all carts.",
		Data:    result,
	}

	response.JSON(w)
}

// GetDummyCartByID gets a cart by id
func GetDummyCartByID(w http.ResponseWriter, r *http.Request) {
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

	result, err := dummyjson.GetDummyCartByID(cartID)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusBadRequest,
			Success:   false,
			Message:   "Error loading cart by id.",
			ErrorCode: errors.DummyJsonError,
		}

		response.JSON(w)
		return
	}

	response := &viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched cart by id.",
		Data:    result,
	}

	response.JSON(w)
}
