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

var myError web.Error
var logRequest web.LogRequest
var logResponse web.LogResponse
var logError web.LogError

type CrmController interface {
	GetContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetContactById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CrmControllerImpl struct {
	crmService qontak_service.CrmService
}

func NewCrmController(crmService qontak_service.CrmService) CrmController {
	return &CrmControllerImpl{
		crmService: crmService,
	}
}

func (controller *CrmControllerImpl) GetContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var getContactRequest qontak_web.CrmGetContactRequest

	helpers.ReadFromRequestBody(request.Body, &getContactRequest)
	contacts, err := controller.crmService.GetContact(request.Context(), getContactRequest)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = contacts
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = getContactRequest

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) GetContactById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var contacts interface{}
	var err error

	contactsId := params.ByName("contact_id")
	if contactsId == "get-params" {
		contacts = controller.crmService.GetParam(request.Context())

	} else {
		contacts, err = controller.crmService.GetContactById(request.Context(), contactsId)
	}

	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = contacts
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

func (controller *CrmControllerImpl) CreateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var crmCreateRequest qontak_web.CrmCreateRequest

	helpers.ReadFromRequestBody(request.Body, &crmCreateRequest)
	contacts, err := controller.crmService.CreateContact(request.Context(), crmCreateRequest)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = contacts
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = crmCreateRequest

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) UpdateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var crmCreateRequest qontak_web.CrmCreateRequest

	contactsId := params.ByName("contact_id")
	helpers.ReadFromRequestBody(request.Body, &crmCreateRequest)
	contacts, err := controller.crmService.UpdateContact(request.Context(), contactsId, crmCreateRequest)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = http.StatusText(http.StatusOK)
		apiResponse.Data = contacts
	}

	logRequest.Url = request.RequestURI
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = crmCreateRequest

	logResponse.Header = writer.Header()
	logResponse.Body = apiResponse
	logResponse.Error = myError.Error

	helpers.WriteLog("info", "event", "incoming request", logRequest, logResponse, logError)
	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) DeleteContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse

	contactsId := params.ByName("contact_id")
	err := controller.crmService.DeleteContact(request.Context(), contactsId)
	if err != nil {
		json.Unmarshal([]byte(err.Error()), &myError)
		writer.WriteHeader(myError.Code)
		apiResponse.Code = myError.Code
		apiResponse.Status = http.StatusText(myError.Code)
		apiResponse.Error = myError.Error
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "deleted"
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
