package qontak_service

import (
	"context"
	"net/http"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"omclabs/go-qontak/app/repositories/qontak_repository"
	"os"
)

type ChatService interface {
	GetIntegrations(ctx context.Context) ([]qontak_web.Channelntegration, error)
	GetIntegrationsByChannel(ctx context.Context, targetChannel string) ([]qontak_web.Channelntegration, error)
	GetWhatsappTemplates(ctx context.Context) ([]qontak_web.WhatsappTemplates, error)
	GetContactList(ctx context.Context) ([]qontak_web.ContactList, error)
	ValidateNumber(ctx context.Context, request qontak_web.ValidateNumberRequest) (qontak_web.ValidateNumberResponse, error)
	SendOtp(ctx context.Context, request qontak_web.SendOtpRequest) (qontak_web.SendMessageResponse, error)
	SendWelcomeMessage(ctx context.Context, welcomeType string, request qontak_web.SendWelcomeMessageRequest) (qontak_web.SendMessageResponse, error)
}

type ChatServiceImpl struct {
	chatRepository qontak_repository.ChatRepository
	client         *http.Client
}

func NewChatService(chatRepository qontak_repository.ChatRepository, client *http.Client) ChatService {
	return &ChatServiceImpl{
		chatRepository: chatRepository,
		client:         client,
	}
}

func (service *ChatServiceImpl) GetIntegrations(ctx context.Context) ([]qontak_web.Channelntegration, error) {
	channelntegration, err := service.chatRepository.GetIntegrations(ctx, service.client)
	if err != nil {
		return channelntegration, err
	}

	return channelntegration, nil
}

func (service *ChatServiceImpl) GetIntegrationsByChannel(ctx context.Context, targetChannel string) ([]qontak_web.Channelntegration, error) {
	channelntegration, err := service.chatRepository.GetIntegrationsByChannel(ctx, service.client, targetChannel)
	if err != nil {
		return channelntegration, err
	}

	return channelntegration, nil
}

func (service *ChatServiceImpl) GetWhatsappTemplates(ctx context.Context) ([]qontak_web.WhatsappTemplates, error) {
	templates, err := service.chatRepository.GetWhatsappTemplates(ctx, service.client)
	if err != nil {
		return templates, err
	}

	return templates, nil
}

func (service *ChatServiceImpl) GetContactList(ctx context.Context) ([]qontak_web.ContactList, error) {
	contactList, err := service.chatRepository.GetContactList(ctx, service.client)
	if err != nil {
		return contactList, err
	}

	return contactList, nil
}

func (service *ChatServiceImpl) ValidateNumber(ctx context.Context, request qontak_web.ValidateNumberRequest) (qontak_web.ValidateNumberResponse, error) {
	validated, err := service.chatRepository.ValidateNumber(ctx, service.client, request)
	if err != nil {
		return validated, err
	}

	return validated, nil
}

func (service *ChatServiceImpl) SendOtp(ctx context.Context, request qontak_web.SendOtpRequest) (qontak_web.SendMessageResponse, error) {

	sendMessage := qontak_web.SendMessageRequest{
		ToNumber:             request.ToNumber,
		ToName:               request.ToName,
		MessageTemplateID:    os.Getenv("QONTAK_CHAT_OTP_TEMPLATE"),
		ChannelIntegrationID: os.Getenv("QONTAK_CHAT_WA_INTEGRATION"),
		Language: struct {
			Code string "json:\"code,omitempty\""
		}{
			Code: "id",
		},
		Parameters: struct {
			Body []struct {
				Key       int    "json:\"key,omitempty\""
				ValueText string "json:\"value_text,omitempty\""
				Value     string "json:\"value,omitempty\""
			} "json:\"body,omitempty\""
			Buttons []struct {
				Index string "json:\"index,omitempty\""
				Type  string "json:\"type,omitempty\""
				Value string "json:\"value,omitempty\""
			} "json:\"buttons,omitempty\""
		}{
			Body: []struct {
				Key       int    "json:\"key,omitempty\""
				ValueText string "json:\"value_text,omitempty\""
				Value     string "json:\"value,omitempty\""
			}{
				{
					Key:       1,
					Value:     "to_name",
					ValueText: request.ToName,
				},
				{
					Key:       2,
					Value:     "otp",
					ValueText: request.Otp,
				},
			},
		},
	}

	validated, err := service.chatRepository.SendMessage(ctx, service.client, sendMessage)
	if err != nil {
		return validated, err
	}

	return validated, nil
}

func (service *ChatServiceImpl) SendWelcomeMessage(ctx context.Context, welcomeType string, request qontak_web.SendWelcomeMessageRequest) (qontak_web.SendMessageResponse, error) {
	var template string

	switch template {
	case "no-store":
		template = os.Getenv("QONTAK_CHAT_WELCOME_NO_STORE_TEMPLATE")
	case "has-store":
		template = os.Getenv("QONTAK_CHAT_WELCOME_HAS_STORE_TEMPLATE")
	default:
		template = ""
	}

	sendMessage := qontak_web.SendMessageRequest{
		ToNumber:             request.ToNumber,
		ToName:               request.ToName,
		MessageTemplateID:    template,
		ChannelIntegrationID: os.Getenv("QONTAK_CHAT_WA_INTEGRATION"),
		Language: struct {
			Code string "json:\"code,omitempty\""
		}{
			Code: "id",
		},
		Parameters: struct {
			Body []struct {
				Key       int    "json:\"key,omitempty\""
				ValueText string "json:\"value_text,omitempty\""
				Value     string "json:\"value,omitempty\""
			} "json:\"body,omitempty\""
			Buttons []struct {
				Index string "json:\"index,omitempty\""
				Type  string "json:\"type,omitempty\""
				Value string "json:\"value,omitempty\""
			} "json:\"buttons,omitempty\""
		}{
			Body: []struct {
				Key       int    "json:\"key,omitempty\""
				ValueText string "json:\"value_text,omitempty\""
				Value     string "json:\"value,omitempty\""
			}{
				{
					Key:       1,
					Value:     "name",
					ValueText: request.Name,
				},
			},
		},
	}

	validated, err := service.chatRepository.SendMessage(ctx, service.client, sendMessage)
	if err != nil {
		return validated, err
	}

	return validated, nil
}
