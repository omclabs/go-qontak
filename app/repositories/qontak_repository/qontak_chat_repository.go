package qontak_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"os"
	"time"
)

type ChatRepository interface {
	GetIntegrations(ctx context.Context, client *http.Client) ([]qontak_web.Channelntegration, error)
	GetIntegrationsByChannel(ctx context.Context, client *http.Client, targetChannel string) ([]qontak_web.Channelntegration, error)
	GetWhatsappTemplates(ctx context.Context, client *http.Client) ([]qontak_web.WhatsappTemplates, error)
	GetContactList(ctx context.Context, client *http.Client) ([]qontak_web.ContactList, error)
	ValidateNumber(ctx context.Context, client *http.Client, request qontak_web.ValidateNumberRequest) (qontak_web.ValidateNumberResponse, error)
}

type ChatRepositoryImpl struct {
}

func NewChatRepository() ChatRepository {
	return &ChatRepositoryImpl{}
}

func (repository *ChatRepositoryImpl) GetIntegrations(ctx context.Context, client *http.Client) ([]qontak_web.Channelntegration, error) {
	var channelIntegration []qontak_web.Channelntegration

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/api/open/v1/integrations"
	method := "GET"
	accessToken, err := getChatToken(client)
	if err != nil {
		return channelIntegration, err
	}

	chatDataRequest, err := sendChatDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return channelIntegration, err
	}

	byteData, err := json.Marshal(chatDataRequest.Data)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return channelIntegration, errors.New(errMessage)
	}

	json.Unmarshal(byteData, &channelIntegration)
	return channelIntegration, nil
}

func (repository *ChatRepositoryImpl) GetIntegrationsByChannel(ctx context.Context, client *http.Client, targetChannel string) ([]qontak_web.Channelntegration, error) {
	var channelIntegration []qontak_web.Channelntegration

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/api/open/v1/integrations?limit=10&target_channel=" + targetChannel
	method := "GET"
	accessToken, err := getChatToken(client)
	if err != nil {
		return channelIntegration, err
	}

	chatDataRequest, err := sendChatDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return channelIntegration, err
	}

	byteData, err := json.Marshal(chatDataRequest.Data)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return channelIntegration, errors.New(errMessage)
	}

	json.Unmarshal(byteData, &channelIntegration)
	return channelIntegration, nil
}

func (repository *ChatRepositoryImpl) GetWhatsappTemplates(ctx context.Context, client *http.Client) ([]qontak_web.WhatsappTemplates, error) {
	var templates []qontak_web.WhatsappTemplates

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/api/open/v1/templates/whatsapp"
	method := "GET"
	accessToken, err := getChatToken(client)
	if err != nil {
		return templates, err
	}

	chatDataRequest, err := sendChatDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return templates, err
	}

	byteData, err := json.Marshal(chatDataRequest.Data)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return templates, errors.New(errMessage)
	}

	json.Unmarshal(byteData, &templates)
	return templates, nil
}

func (repository *ChatRepositoryImpl) GetContactList(ctx context.Context, client *http.Client) ([]qontak_web.ContactList, error) {
	var contactList []qontak_web.ContactList

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/api/open/v1/contacts/contact_lists"
	method := "GET"
	accessToken, err := getChatToken(client)
	if err != nil {
		return contactList, err
	}

	chatDataRequest, err := sendChatDataRequest(url, method, accessToken, nil, client)
	if err != nil {
		return contactList, err
	}

	byteData, err := json.Marshal(chatDataRequest.Data)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return contactList, errors.New(errMessage)
	}

	json.Unmarshal(byteData, &contactList)
	return contactList, nil
}

func (repository *ChatRepositoryImpl) ValidateNumber(ctx context.Context, client *http.Client, request qontak_web.ValidateNumberRequest) (qontak_web.ValidateNumberResponse, error) {
	var validated qontak_web.ValidateNumberResponse

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/api/open/v1/broadcasts/contacts"
	method := "POST"
	accessToken, err := getChatToken(client)
	if err != nil {
		return validated, err
	}

	chatDataRequest, err := sendChatDataRequest(url, method, accessToken, request, client)
	if err != nil {
		return validated, err
	}

	byteData, err := json.Marshal(chatDataRequest.Data)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return validated, errors.New(errMessage)
	}

	json.Unmarshal(byteData, &validated)
	return validated, nil
}

func getChatToken(client *http.Client) (string, error) {
	fileName := "./files/chat_auth.json"
	var token string = ""
	if helpers.CheckDirOrFileExists(fileName) {
		file := helpers.ReadJsonFile(fileName)
		jsonData := qontak_web.CrmAuth{}

		json.Unmarshal(file, &jsonData)
		expiresIn := jsonData.CreatedAt + jsonData.ExpiresIn
		timeNow := time.Now().Unix()

		if timeNow > int64(expiresIn) {
			accessToken, err := authChatUser(client)
			if err != nil {
				return accessToken, err
			}
			token = accessToken
		} else {
			token = jsonData.AccessToken
		}
	} else {
		accessToken, err := authChatUser(client)
		if err != nil {
			return accessToken, err
		}
		token = accessToken
	}

	return token, nil
}

func authChatUser(client *http.Client) (string, error) {
	accessToken := ""

	url := os.Getenv("QONTAK_CHAT_BASE_URL") + "/oauth/token"
	method := "POST"
	payload := fmt.Sprintf(`{"grant_type": "%s","username": "%s","password": "%s","client_id": "%s","client_secret": "%s"}`,
		os.Getenv("QONTAK_CHAT_GRANT_TYPE"),
		os.Getenv("QONTAK_CHAT_USERNAME"),
		os.Getenv("QONTAK_CHAT_PASSWORD"),
		os.Getenv("QONTAK_CHAT_CLIENT_ID"),
		os.Getenv("QONTAK_CHAT_CLIENT_SECRET"))

	result, err := sendAuthRequest(url, method, payload, client)
	if err != nil {
		return accessToken, err
	}

	jsonAuth, err := json.Marshal(result)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return accessToken, errors.New(errMessage)
	}

	helpers.WriteJsonFile("files", "chat_auth.json", true, string(jsonAuth))
	accessToken = result.AccessToken
	return accessToken, nil
}

func sendChatDataRequest(url string, method string, token string, payload interface{}, client *http.Client) (qontak_web.ChatDataRequest, error) {
	var chatDataRequest qontak_web.ChatDataRequest
	var body bytes.Buffer

	json.NewEncoder(&body).Encode(payload)
	request, err := http.NewRequest(method, url, &body)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return chatDataRequest, errors.New(errMessage)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		errMessage = helpers.MakeErrorMessage(500, err.Error())
		return chatDataRequest, errors.New(errMessage)
	}

	defer response.Body.Close()
	chatDataRequest.Status = response.Status
	json.NewDecoder(response.Body).Decode(&chatDataRequest)

	logRequest.Url = url
	logRequest.Header = request.Header
	logRequest.Method = request.Method
	logRequest.Payload = body

	logResponse.Header = response.Header
	logResponse.Body = chatDataRequest

	helpers.WriteLog("info", "external", "calling external service", logRequest, logResponse, logError)
	if response.StatusCode != 200 && response.StatusCode != 201 {
		switch response.StatusCode {
		case 401:
			errMessage = helpers.MakeErrorMessage(401, chatDataRequest.Error.Messages)
			return chatDataRequest, errors.New(errMessage)
		default:
			errMessage = helpers.MakeErrorMessage(response.StatusCode, chatDataRequest.Error.Messages)
			return chatDataRequest, errors.New(errMessage)
		}
	}

	return chatDataRequest, nil
}
