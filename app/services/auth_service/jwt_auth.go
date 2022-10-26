package auth_service

import (
	"context"
	"errors"
	"fmt"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web/auth"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtAuthService interface {
	AuthUser(ctx context.Context, request auth.AuthRequest) (auth.AuthResponse, error)
}

type JwtAuthServiceImpl struct {
}

func NewJwtAuthService() JwtAuthService {
	return &JwtAuthServiceImpl{}
}

func (service *JwtAuthServiceImpl) AuthUser(ctx context.Context, request auth.AuthRequest) (auth.AuthResponse, error) {
	var response auth.AuthResponse
	signature := os.Getenv("TOKEN_SECRET")
	expIn := time.Now().Unix() + 60*60*24*365 // 1 year

	if request.Username != os.Getenv("DEFAULT_USERNAME") && request.Password != os.Getenv("DEFAULT_PASSWORD") {
		errMessage := helpers.MakeErrorMessage(401, "failed to authenticate user")
		return response, errors.New(errMessage)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  expIn,
		"user": request.Username,
	})

	accessToken, err := token.SignedString([]byte(signature))
	if err != nil {
		errMessage := helpers.MakeErrorMessage(500, err.Error())
		return response, errors.New(errMessage)
	}

	response.AccessToken = accessToken
	response.CreatedAt = time.Now().Unix()
	response.ExpiredAt = int64(expIn)

	fmt.Println(response)
	return response, nil
}
