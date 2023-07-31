package service

import (
	"crypto/rand"
	"api-key-service/entity"
	"api-key-service/repository"
)

const KeyByteSize int = 32

type ApiKeyService interface {
	// Generates key for owner
	GenerateKey(string) (*entity.ApiKey, error)

	// Find keys for owner
	FindKeys(string) (*[]entity.ApiKey, error)

	// Authenticates with internal key bytes
	Authenticate([]byte) (bool, error)
}

type Service struct {
	repository repository.ApiKeyRepository
}

func NewService(repository repository.ApiKeyRepository) ApiKeyService {
	return &Service{
		repository: repository,
	}
}

func (service *Service) GenerateKey(owner string) (*entity.ApiKey, error) {
	// Cryptographically secure generation of 32 bytes
	bytes := make([]byte, KeyByteSize)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	key := &entity.ApiKey{
		Owner: owner,
		Key:   bytes,
	}

	err = service.repository.Save(key)
	return key, err
}

func (service *Service) FindKeys(owner string) (*[]entity.ApiKey, error) {
	return service.repository.FindByOwner(owner)
}

func (service *Service) Authenticate(key []byte) (bool, error) {
	apiKey, err := service.repository.FindByKey(key)
	return apiKey != nil, err
}
