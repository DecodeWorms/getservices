package migrator

import (
	"context"

	"github.com/DecodeWorms/getservices/models"

	"gorm.io/gorm"
)

type ServiceProviderMigrations interface {
	Create(ctx context.Context) error
	Address(ctx context.Context) error
}

type ServiceProvider struct {
	Client *gorm.DB
}

func NewServeProvider(cl *gorm.DB) ServiceProvider {
	return ServiceProvider{
		Client: cl,
	}
}

func (s ServiceProvider) Create(ctx context.Context) error {
	var c models.ServiceProvider
	return s.Client.AutoMigrate(&c)
}

func (s ServiceProvider) Address(ctx context.Context) error {
	var ad models.ServiceProviderAddress
	return s.Client.AutoMigrate(&ad)
}
