package controller

import (
	"api-key-service/entity"
	"api-key-service/service"
)

type ApiKeyController interface {
	// Key generation endpoint
	GenerateKey(entity.GenerateKeyRequest) (*entity.ApiKey, error)

	// Key lookup endpoint
	FindKeys(entity.FindKeysRequest) (*[]entity.ApiKey, error)

	// Authentication endpoint
	Authenticate(entity.AuthenticateRequest) (bool, error)
}

type Controller struct {
	Service service.ApiKeyService
}

func NewController(service *service.ApiKeyService) ApiKeyController {
	return &Controller{Service: *service}
}

func (controller *Controller) GenerateKey(request entity.GenerateKeyRequest) (*entity.ApiKey, error) {
	owner := request.Owner
	return controller.Service.GenerateKey(owner)
}

func (controller *Controller) FindKeys(request entity.FindKeysRequest) (*[]entity.ApiKey, error) {
	owner := request.Owner
	return controller.Service.FindKeys(owner)
}

func (controller *Controller) Authenticate(request entity.AuthenticateRequest) (bool, error) {
	key := request.Key
	return controller.Service.Authenticate(key)
}
