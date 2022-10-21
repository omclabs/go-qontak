package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadFromRequestBody(requestBody io.ReadCloser, result interface{}) {
	decoder := json.NewDecoder(requestBody)
	decoder.Decode(result)
	// PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
	// PanicIfError(err)
}

func MakeLogrusFields(origin interface{}) map[string]interface{} {
	var objResult map[string]interface{}
	objJson, _ := json.Marshal(origin)
	// PanicIfError(err)
	json.Unmarshal(objJson, &objResult)
	return objResult
}

func MakeErrorMessage(code int, message string) string {
	return fmt.Sprintf(`{"code":%d, "error":"%s"}`, code, message)
}
