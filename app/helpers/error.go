package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"omclabs/go-qontak/app/models/web"
	"runtime"
)

var myError web.Error
var logError web.LogError
var logRequest web.LogRequest
var logResponse web.LogResponse

// func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
// 	exception, ok := err.(validator.ValidationErrors)
// 	if ok {
// 		writer.Header().Set("Content-Type", "application/json")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		apiResponse := web.ApiResponse{
// 			Code:   http.StatusBadRequest,
// 			Status: "Bad Request",
// 			Error:  exception.Error(),
// 		}
// 		WriteToResponseBody(writer, apiResponse)
// 		return true
// 	} else {
// 		return false
// 	}
// }

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	json.Unmarshal([]byte(fmt.Sprintf("%v", err)), &myError)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(myError.Code)

	apiResponse := web.ApiResponse{
		Code:   myError.Code,
		Status: http.StatusText(myError.Code),
		Error:  myError.Error,
	}

	_, file, line, _ := runtime.Caller(1)
	logError.Caller = file + " line " + fmt.Sprint(line)
	logError.ErrorDesc = myError.Error

	WriteLog("error", "error", "internal server error", logRequest, logResponse, logError)
	WriteToResponseBody(writer, apiResponse)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
		// panic(SetError(500, err.Error()))
	}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// if validationErrors(writer, request, err) {
	// 	return
	// }

	internalServerError(writer, request, err)
}
