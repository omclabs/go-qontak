package qontak_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"os"
	"strconv"
	"strings"
	"time"
)

type CrmRepository interface {
	GetParam(ctx context.Context) qontak_web.CrmParams
	GetContact(ctx context.Context, request qontak_web.CrmGetContactRequest, client *http.Client) ([]qontak_web.CrmContacts, error)
	GetContactById(ctx context.Context, id int, client *http.Client) (qontak_web.CrmContacts, error)
	CreateContact(ctx context.Context, request qontak_web.CrmCreateRequest, client *http.Client) (qontak_web.CrmContacts, error)
	UpdateContact(ctx context.Context, id int, request qontak_web.CrmCreateRequest, client *http.Client) error
	DeleteContact(ctx context.Context, id int, client *http.Client) error
}

type CrmRepositoryImpl struct {
}

func NewCrmRepository() CrmRepository {
	return &CrmRepositoryImpl{}
}

func (repository *CrmRepositoryImpl) GetParam(ctx context.Context) qontak_web.CrmParams {
	files := helpers.ReadJsonFile("./files/crm_param.json")
	crmParams := qontak_web.CrmParams{}
	json.Unmarshal(files, &crmParams)

	return crmParams
}

func (repository *CrmRepositoryImpl) GetContact(ctx context.Context, request qontak_web.CrmGetContactRequest, client *http.Client) ([]qontak_web.CrmContacts, error) {
	var crmContacts []qontak_web.CrmContacts

	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/api/v3.1/contacts"
	method := "GET"
	accessToken, err := getToken(client)
	if err != nil {
		return crmContacts, err
	}

	params := "?name=" + request.Name + "&email=" + request.Email + "&phone=" + request.Phone
	url += params

	crmContactData, err := sendDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return crmContacts, err
	}

	byteData, err := json.Marshal(crmContactData.Response)
	if err != nil {
		return crmContacts, helpers.GetErrInternalServer()
	}

	json.Unmarshal(byteData, &crmContacts)
	return crmContacts, nil
}

func (repository *CrmRepositoryImpl) GetContactById(ctx context.Context, id int, client *http.Client) (qontak_web.CrmContacts, error) {
	var crmContacts qontak_web.CrmContacts

	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/api/v3.1/contacts/" + strconv.Itoa(id)
	method := "GET"
	accessToken, err := getToken(client)
	if err != nil {
		return crmContacts, err
	}

	crmDataRequest, err := sendDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return crmContacts, err
	}

	byteData, err := json.Marshal(crmDataRequest.Response)
	if err != nil {
		return crmContacts, helpers.GetErrInternalServer()
	}

	json.Unmarshal(byteData, &crmContacts)
	return crmContacts, nil
}

func (repository *CrmRepositoryImpl) CreateContact(ctx context.Context, request qontak_web.CrmCreateRequest, client *http.Client) (qontak_web.CrmContacts, error) {
	var crmContacts qontak_web.CrmContacts
	var payload bytes.Buffer

	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/api/v3.1/contacts"
	method := "POST"
	accessToken, err := getToken(client)
	if err != nil {
		return crmContacts, err
	}

	err = json.NewEncoder(&payload).Encode(request)
	if err != nil {
		return crmContacts, helpers.GetErrInternalServer()
	}

	CrmContactsData, err := sendDataRequest(url, method, accessToken, &payload, client)
	if err != nil {
		return crmContacts, err
	}

	byteData, err := json.Marshal(CrmContactsData.Response)
	if err != nil {
		return crmContacts, helpers.GetErrInternalServer()
	}

	json.Unmarshal(byteData, &crmContacts)
	return crmContacts, nil
}

func (repository *CrmRepositoryImpl) UpdateContact(ctx context.Context, id int, request qontak_web.CrmCreateRequest, client *http.Client) error {
	var payload bytes.Buffer

	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/api/v3.1/contacts/" + strconv.Itoa(id)
	method := "PUT"
	accessToken, err := getToken(client)
	if err != nil {
		return err
	}

	err = json.NewEncoder(&payload).Encode(request)
	if err != nil {
		return helpers.GetErrInternalServer()
	}

	_, err = sendDataRequest(url, method, accessToken, &payload, client)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CrmRepositoryImpl) DeleteContact(ctx context.Context, id int, client *http.Client) error {
	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/api/v3.1/contacts/" + strconv.Itoa(id)
	method := "DELETE"
	accessToken, err := getToken(client)
	if err != nil {
		return err
	}

	_, err = sendDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return err
	}

	return nil
}

func getToken(client *http.Client) (string, error) {
	fileName := "./files/crm_auth.json"
	var token string = ""
	if helpers.CheckJsonFileExists(fileName) {
		file := helpers.ReadJsonFile(fileName)
		jsonData := qontak_web.CrmAuth{}

		json.Unmarshal(file, &jsonData)
		expiresIn := jsonData.CreatedAt + jsonData.ExpiresIn
		timeNow := time.Now().Unix()

		if timeNow > int64(expiresIn) {
			accessToken, err := authUser(client)
			if err != nil {
				return accessToken, err
			}
			token = accessToken
		} else {
			token = jsonData.AccessToken
		}
	} else {
		accessToken, err := authUser(client)
		if err != nil {
			return accessToken, err
		}
		token = accessToken
	}

	return token, nil
}

func authUser(client *http.Client) (string, error) {
	accessToken := ""
	url := os.Getenv("QONTAK_CRM_BASE_URL") + "/oauth/token"
	method := "POST"
	jsonString := fmt.Sprintf(`{"grant_type": "%s","username": "%s","password": "%s"}`,
		os.Getenv("QONTAK_CRM_GRANT_TYPE"), os.Getenv("QONTAK_CRM_USERNAME"), os.Getenv("QONTAK_CRM_PASSWORD"))
	payload := strings.NewReader(jsonString)
	result, err := sendAuthRequest(url, method, payload, client)
	if err != nil {
		return accessToken, err
	}

	jsonAuth, err := json.Marshal(result)
	if err != nil {
		return accessToken, helpers.GetErrInternalServer()
	}

	helpers.WriteJsonFile("files", "crm_auth.json", true, string(jsonAuth))
	accessToken = result.AccessToken
	return accessToken, nil
}

func sendAuthRequest(url string, method string, payload io.Reader, client *http.Client) (qontak_web.CrmAuth, error) {
	var crmAuth qontak_web.CrmAuth
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return crmAuth, helpers.GetErrInternalServer()
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return crmAuth, helpers.GetErrInternalServer()
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return crmAuth, helpers.MapHttpErrorCode(int(response.StatusCode))
	}

	crmAuth.Code = response.StatusCode
	crmAuth.Status = response.Status
	json.NewDecoder(response.Body).Decode(&crmAuth)
	return crmAuth, nil
}

func sendDataRequest(url string, method string, token string, payload io.Reader, client *http.Client) (qontak_web.CrmContactsData, error) {
	var crmDataRequest qontak_web.CrmContactsData
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return crmDataRequest, helpers.GetErrInternalServer()
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
	response, err := client.Do(request)
	if err != nil {
		return crmDataRequest, helpers.GetErrInternalServer()
	}

	defer response.Body.Close()
	crmDataRequest.Code = response.StatusCode
	crmDataRequest.Status = response.Status
	json.NewDecoder(response.Body).Decode(&crmDataRequest)
	if response.StatusCode != 200 {
		return crmDataRequest, helpers.MapHttpErrorCode(int(response.StatusCode))
	}

	if crmDataRequest.Meta.Status != 200 {
		return crmDataRequest, helpers.MapHttpErrorCode(int(crmDataRequest.Meta.Status))
	}

	return crmDataRequest, nil
}
