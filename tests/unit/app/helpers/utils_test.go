package helpers_test

import (
	"fmt"
	"io/ioutil"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeErrorMessage(t *testing.T) {
	code := 500
	message := "internal server error"
	expected := fmt.Sprintf(`{"code":%d, "error":"%s"}`, code, message)
	result := helpers.MakeErrorMessage(500, "internal server error")

	assert.Equal(t, expected, result, "expected value not equal to result value")
}

func TestMakeLogrusFields(t *testing.T) {
	var logrusError web.LogError
	var logRequest web.LogRequest
	var logResponse web.LogResponse

	logrusFields := web.LoggerData{
		Error:    logrusError,
		Request:  logRequest,
		Response: logResponse,
	}

	result := helpers.MakeLogrusFields(logrusFields)
	assert.Contains(t, result, "request")
	assert.Contains(t, result, "error")
	assert.Contains(t, result, "response")
}

func TestReadFromRequestBody(t *testing.T) {
	var objInterface struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	objReader := ioutil.NopCloser(strings.NewReader(`{"username":"test", "password":"rahasia", "token":"adadeh"}`))
	helpers.ReadFromRequestBody(objReader, &objInterface)
	assert.Equal(t, objInterface.Username, "test")
	assert.Equal(t, objInterface.Password, "rahasia")
	assert.Equal(t, objInterface.Token, "adadeh")
}
