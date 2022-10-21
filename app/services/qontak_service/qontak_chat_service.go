package qontak_service

import (
	"context"
	"net/http"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"omclabs/go-qontak/app/repositories/qontak_repository"
)

type ChatService interface {
	GetIntegrations(ctx context.Context) ([]qontak_web.Channelntegration, error)
	GetIntegrationsByChannel(ctx context.Context, targetChannel string) ([]qontak_web.Channelntegration, error)
	GetWhatsappTemplates(ctx context.Context) ([]qontak_web.WhatsappTemplates, error)
	GetContactList(ctx context.Context) ([]qontak_web.ContactList, error)
	ValidateNumber(ctx context.Context, request qontak_web.ValidateNumberRequest) (qontak_web.ValidateNumberResponse, error)
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
