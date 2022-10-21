package qontak_repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"strings"
)

var errMessage string
var logError web.LogError
var logRequest web.LogRequest
var logResponse web.LogResponse

func sendAuthRequest(url string, method string, payload interface{}, client *http.Client) (qontak_web.CrmAuth, error) {
	var crmAuth qontak_web.CrmAuth

	body := strings.NewReader(fmt.Sprintf("%v", payload))
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return crmAuth, errors.New(errMessage)
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return crmAuth, errors.New(errMessage)
	}

	crmAuth.Code = response.StatusCode
	crmAuth.Status = response.Status
	json.NewDecoder(response.Body).Decode(&crmAuth)
	defer response.Body.Close()

	logRequest.Url = url
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = payload

	logResponse.Header = response.Header
	logResponse.Body = crmAuth

	helpers.WriteLog("info", "external", "calling external service", logRequest, logResponse, logError)
	if response.StatusCode != 200 && response.StatusCode != 201 {
		errMessage := helpers.MakeErrorMessage(response.StatusCode, response.Status)
		return crmAuth, errors.New(errMessage)
	}

	return crmAuth, nil
}
