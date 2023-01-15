package migrator

import (
	"context"
	"getservices/models"

	"gorm.io/gorm"
)

type ClientMigrations interface {
	Create(ctx context.Context) error
	Address(ctx context.Context) error
}

type Client struct {
	db *gorm.DB
}

func NewCleint(db *gorm.DB) Client {
	return Client{
		db: db,
	}

}

func (c Client) Create(ctx context.Context) error {
	var d models.Client
	return c.db.AutoMigrate(&d)
}

func (c Client) Address(ctx context.Context) error {
	var ad models.Address
	return c.db.AutoMigrate(&ad)
}
