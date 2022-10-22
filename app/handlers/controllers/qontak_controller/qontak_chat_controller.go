package qontak_controller

import (
	"encoding/json"
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"omclabs/go-qontak/app/services/qontak_service"

	"github.com/julienschmidt/httprouter"
)

type ChatController interface {
	GetIntegrations(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetIntegrationsByChannel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetWhatsappIntegration(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetWhatsappTemplates(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetContactList(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ValidateNumber(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SendWelcomeMessage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ChatControllerImpl struct {
	chatService qontak_service.ChatService
}

func NewChatController(chatService qontak_service.ChatService) ChatController {
	return &ChatControllerImpl{
		chatService: chatService,
	}
}

func (controller *ChatControllerImpl) GetIntegrations(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	integrations, err := controller.chatService.GetIntegrations(request.Context())
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = integrations
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) GetIntegrationsByChannel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	channel := params.ByName("channel")
	integrations, err := controller.chatService.GetIntegrationsByChannel(request.Context(), channel)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = integrations
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) GetWhatsappIntegration(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	integrations, err := controller.chatService.GetIntegrationsByChannel(request.Context(), "wa")
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = integrations
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) GetWhatsappTemplates(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	templates, err := controller.chatService.GetWhatsappTemplates(request.Context())
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = templates
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) GetContactList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	contactList, err := controller.chatService.GetContactList(request.Context())
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = contactList
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) ValidateNumber(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var validateNumberRequest qontak_web.ValidateNumberRequest

	helpers.ReadFromRequestBody(request.Body, &validateNumberRequest)

	validated, err := controller.chatService.ValidateNumber(request.Context(), validateNumberRequest)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = validated
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var otp qontak_web.SendOtpRequest

	helpers.ReadFromRequestBody(request.Body, &otp)

	validated, err := controller.chatService.SendOtp(request.Context(), otp)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = validated
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *ChatControllerImpl) SendWelcomeMessage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var welcome qontak_web.SendWelcomeMessageRequest

	welcomeType := params.ByName("type")
	helpers.ReadFromRequestBody(request.Body, &welcome)
	validated, err := controller.chatService.SendWelcomeMessage(request.Context(), welcomeType, welcome)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = validated
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}
