package qontak_controller

import (
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"omclabs/go-qontak/app/services/qontak_service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CrmController interface {
	GetParam(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
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

func (controller *CrmControllerImpl) GetParam(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	crmParams := controller.crmService.GetParam(request.Context())

	apiResponse := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   crmParams,
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) GetContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var getContactRequest qontak_web.CrmGetContactRequest
	helpers.ReadFromRequestBody(request, &getContactRequest)

	contacts, err := controller.crmService.GetContact(request.Context(), getContactRequest)

	if err != nil {
		apiResponse.Code = helpers.MapHttpStatusCode(err)
		apiResponse.Status = err.Error()
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "ok"
		apiResponse.Data = contacts
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) GetContactById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	contactsId := params.ByName("id")
	id, err := strconv.Atoi(contactsId)
	helpers.PanicIfError(err)

	contacts, err := controller.crmService.GetContactById(request.Context(), id)
	if err != nil {
		apiResponse.Code = helpers.MapHttpStatusCode(err)
		apiResponse.Status = err.Error()
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "ok"
		apiResponse.Data = contacts
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) CreateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var crmCreateRequest qontak_web.CrmCreateRequest

	helpers.ReadFromRequestBody(request, &crmCreateRequest)
	contacts, err := controller.crmService.CreateContact(request.Context(), crmCreateRequest)

	if err != nil {
		apiResponse.Code = helpers.MapHttpStatusCode(err)
		apiResponse.Status = err.Error()
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "ok"
		apiResponse.Data = contacts
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) UpdateContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	var crmCreateRequest qontak_web.CrmCreateRequest

	contactsId := params.ByName("id")
	id, err := strconv.Atoi(contactsId)
	helpers.PanicIfError(err)

	helpers.ReadFromRequestBody(request, &crmCreateRequest)
	contacts, err := controller.crmService.UpdateContact(request.Context(), id, crmCreateRequest)

	if err != nil {
		apiResponse.Code = helpers.MapHttpStatusCode(err)
		apiResponse.Status = err.Error()
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "ok"
		apiResponse.Data = contacts
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}

func (controller *CrmControllerImpl) DeleteContact(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var apiResponse web.ApiResponse
	contactsId := params.ByName("id")
	id, err := strconv.Atoi(contactsId)
	helpers.PanicIfError(err)

	err = controller.crmService.DeleteContact(request.Context(), id)
	if err != nil {
		apiResponse.Code = helpers.MapHttpStatusCode(err)
		apiResponse.Status = err.Error()
	} else {
		apiResponse.Code = http.StatusOK
		apiResponse.Status = "deleted"
	}

	helpers.WriteToResponseBody(writer, apiResponse)
}
