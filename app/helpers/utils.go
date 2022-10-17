package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func StructToMap(origin interface{}) map[string]interface{} {
	var objResult map[string]interface{}
	objJson, err := json.Marshal(origin)
	PanicIfError(err)
	json.Unmarshal(objJson, &objResult)
	return objResult
}

func MakeErrorMessage(code int, message string) string {
	type myError struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}

	objError := myError{
		Code:  code,
		Error: message,
	}

	objJson, err := json.Marshal(objError)
	PanicIfError(err)

	return string(objJson)
}

func JsonToInterface(origin string, target interface{}) {
	err := json.Unmarshal([]byte(origin), &target)
	PanicIfError(err)
}
