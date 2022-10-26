package auth_controller

import (
	"encoding/json"
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"omclabs/go-qontak/app/models/web/auth"
	"omclabs/go-qontak/app/services/auth_service"

	"github.com/julienschmidt/httprouter"
)

var myError web.Error
var logRequest web.LogRequest
var logResponse web.LogResponse
var logError web.LogError

type JwtAuthController interface {
	AuthUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type JwtAuthControllerImpl struct {
	authService auth_service.JwtAuthService
}

func NewJwtAuthController(jwtAuth auth_service.JwtAuthService) JwtAuthController {
	return &JwtAuthControllerImpl{
		authService: jwtAuth,
	}
}

func (controller *JwtAuthControllerImpl) AuthUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var authRequest auth.AuthRequest

	helpers.ReadFromRequestBody(request.Body, &authRequest)
	authResponse, err := controller.authService.AuthUser(request.Context(), authRequest)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = authResponse
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = authRequest

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}
