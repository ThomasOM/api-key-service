package repository

import (
	"errors"
	"api-key-service/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiKeyRepository interface {
	// Saves an API Key to the database
	Save(*entity.ApiKey) error

	// Finds all API Keys by from an owner
	FindByOwner(string) (*[]entity.ApiKey, error)

	// Finds the API Key by the internal key bytes
	FindByKey([]byte) (*entity.ApiKey, error)
}

type Repository struct {
	database *gorm.DB
}

// Creates a new repository implementation with a database
func NewRepository() ApiKeyRepository {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entity.ApiKey{})

	return &Repository{
		database: database,
	}
}

func (repository *Repository) Save(key *entity.ApiKey) error {
	return repository.database.Create(key).Error
}

func (repository *Repository) FindByOwner(owner string) (*[]entity.ApiKey, error) {
	var keys []entity.ApiKey
	err := repository.database.Where("owner = ?", owner).Find(&keys).Error

	// Do not handle as error sinc there is just no record
	if errors.Is(err, gorm.ErrRecordNotFound) {
		keys = []entity.ApiKey{}
		err = nil
	}

	return &keys, err
}

func (repository *Repository) FindByKey(key []byte) (*entity.ApiKey, error) {
	var apiKey entity.ApiKey
	err := repository.database.Where("key = ?", key).First(&apiKey).Error

	// Do not handle as error sinc there is just no record
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &apiKey, err
}
