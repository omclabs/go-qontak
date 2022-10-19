package qontak_service

import (
	"context"
	"net/http"
	"omclabs/go-qontak/app/models/web/qontak_web"
	"omclabs/go-qontak/app/repositories/qontak_repository"
)

type CrmService interface {
	GetParam(ctx context.Context) qontak_web.CrmParams
	GetContact(ctx context.Context, request qontak_web.CrmGetContactRequest) ([]qontak_web.CrmContacts, error)
	GetContactById(ctx context.Context, id string) (qontak_web.CrmContacts, error)
	CreateContact(ctx context.Context, request qontak_web.CrmCreateRequest) (qontak_web.CrmContacts, error)
	UpdateContact(ctx context.Context, id string, request qontak_web.CrmCreateRequest) (qontak_web.CrmContacts, error)
	DeleteContact(ctx context.Context, id string) error
}

type CrmServiceImpl struct {
	crmRepository qontak_repository.CrmRepository
	client        *http.Client
}

func NewCrmService(crmRepository qontak_repository.CrmRepository, client *http.Client) CrmService {
	return &CrmServiceImpl{
		crmRepository: crmRepository,
		client:        client,
	}
}

func (service *CrmServiceImpl) GetParam(ctx context.Context) qontak_web.CrmParams {
	crmParams := service.crmRepository.GetParam(ctx)

	return qontak_web.CrmParams(crmParams)
}

func (service *CrmServiceImpl) GetContact(ctx context.Context, request qontak_web.CrmGetContactRequest) ([]qontak_web.CrmContacts, error) {
	crmContacts, err := service.crmRepository.GetContact(ctx, request, service.client)
	if err != nil {
		return crmContacts, err
	}

	return crmContacts, nil
}

func (service *CrmServiceImpl) GetContactById(ctx context.Context, id string) (qontak_web.CrmContacts, error) {
	var crmContacts qontak_web.CrmContacts
	crmContacts, err := service.crmRepository.GetContactById(ctx, id, service.client)
	if err != nil {
		return crmContacts, err
	}

	return crmContacts, nil
}

func (service *CrmServiceImpl) CreateContact(ctx context.Context, request qontak_web.CrmCreateRequest) (qontak_web.CrmContacts, error) {
	var crmContacts qontak_web.CrmContacts
	crmContacts, err := service.crmRepository.CreateContact(ctx, request, service.client)
	if err != nil {
		return crmContacts, err
	}

	return crmContacts, nil
}

func (service *CrmServiceImpl) UpdateContact(ctx context.Context, id string, request qontak_web.CrmCreateRequest) (qontak_web.CrmContacts, error) {
	var crmContacts qontak_web.CrmContacts
	err := service.crmRepository.UpdateContact(ctx, id, request, service.client)
	if err != nil {
		return crmContacts, err
	}
	crmContacts, err = service.crmRepository.GetContactById(ctx, id, service.client)
	if err != nil {
		return crmContacts, err
	}
	return crmContacts, nil
}

func (service *CrmServiceImpl) DeleteContact(ctx context.Context, id string) error {
	err := service.crmRepository.DeleteContact(ctx, id, service.client)
	if err != nil {
		return err
	}

	return nil
}
