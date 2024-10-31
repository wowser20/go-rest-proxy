package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"go-rest-proxy/internal/errors"
	"go-rest-proxy/internal/viewmodels"
	"go-rest-proxy/module/iam/handler/types"
)

// GenerateToken generates a jwt token
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	// create access token
	accessTokenClaims := jwt.MapClaims{
		"iss": "rest-proxy",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "Error generating token.",
			ErrorCode: errors.ServerError,
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusCreated,
		Success: true,
		Message: "Successfully generated the token.",
		Data: &types.GenerateTokenResponse{
			AccessToken: token,
		},
	}

	response.JSON(w)
}
