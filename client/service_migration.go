package client

import (
	"context"
	"getservices/models"

	"gorm.io/gorm"
)

type ServiceMigrations interface {
	Create(ctx context.Context) error
	Address(ctx context.Context) error
}

type Service struct {
	db *gorm.DB
}

func NewServices(db *gorm.DB) Service {
	return Service{
		db: db,
	}
}

func (s Service) Create(ctx context.Context) error {
	var service models.Services
	return s.db.AutoMigrate(&service)
}

func (s Service) Address(ctx context.Context) error {
	var serviceAddress models.ServiceAddress
	return s.db.AutoMigrate(&serviceAddress)
}
