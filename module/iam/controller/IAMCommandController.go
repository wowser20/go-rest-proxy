package handler

import (
	"net/http"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/module/iam/controller/types"
	"go-rest-proxy/module/iam/service"
)

type IAMCommandController struct {
	service.IAMCommandService
}

// GenerateToken generates a jwt token
func (controller *IAMCommandController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	res, err := controller.IAMCommandService.GenerateToken()
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

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusCreated,
		Success: true,
		Message: "Successfully generated the token.",
		Data: &types.GenerateTokenResponse{
			AccessToken: res,
		},
	}

	response.JSON(w)
}
