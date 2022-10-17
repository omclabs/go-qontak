package helpers

import (
	"net/http"
	"omclabs/go-qontak/app/models/web"

	"github.com/go-playground/validator/v10"
)

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		apiResponse := web.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  exception.Error(),
		}
		WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	apiResponse := web.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Error:  err,
	}
	WriteToResponseBody(writer, apiResponse)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
		// panic(SetError(500, err.Error()))
	}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}
