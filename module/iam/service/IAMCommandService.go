package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	apiError "go-rest-proxy/internal/errors"
)

type IAMCommandService struct {
}

// GenerateToken generate token
func (service *IAMCommandService) GenerateToken() (string, error) {
	// create access token
	accessTokenClaims := jwt.MapClaims{
		"iss": "rest-proxy",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(), // 15 minutes expiration
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Print(err)
		return "", errors.New(apiError.UnauthorizedAccess)
	}

	return token, nil
}
