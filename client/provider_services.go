package client

import (
	"getservices/models"
	"time"

	"gorm.io/gorm"
)

type ProviderServices interface {
	Create(cl models.ServiceProvider) error
	Provider(providerId string) (models.ServiceProvider, error)
	Providers() ([]models.ServiceProvider, error)
	ProviderByEmail(email string) (models.ServiceProvider, error)
	ProviderByPhoneNumber(phoneNumber string) (models.ServiceProvider, error)
	Update(providerId string, cl models.ServiceProvider) error
	DeactivateAccount(clientId string) error
	ActivateAccount(email string) (models.ServiceProvider, error)
}

type ServiceProviderAccount struct {
	db *gorm.DB
}

func NewServiceProviderAccount(db *gorm.DB) ServiceProviderAccount {
	return ServiceProviderAccount{
		db: db,
	}
}

func (provider ServiceProviderAccount) Create(pr models.ServiceProvider) error {
	pr.CreatedAt = time.Now()
	pr.UpdatedAt = time.Now()

	data := models.ServiceProvider{
		ServiceProviderId:       pr.ServiceProviderId,
		ServiceProviderIdentity: pr.ServiceProviderIdentity,
		FirstName:               pr.FirstName,
		LastName:                pr.LastName,
		PhoneNumber:             pr.PhoneNumber,
		Email:                   pr.Email,
		Password:                pr.Password,
	}
	return provider.db.Create(data).Error
}

func (provider ServiceProviderAccount) Provider(providerId string) (models.ServiceProvider, error) {
	var p models.ServiceProvider
	return p, provider.db.Where("client_id = ?", providerId).First(&p).Error
}

func (provider ServiceProviderAccount) Providers() ([]models.ServiceProvider, error) {
	var p []models.ServiceProvider
	return p, provider.db.Find(&p).Error
}

func (provider ServiceProviderAccount) ProviderByEmail(email string) (models.ServiceProvider, error) {
	var p models.ServiceProvider
	return p, provider.db.Where("email = ?", email).First(&p).Error
}

func (provider ServiceProviderAccount) ClientByPhoneNumber(phoneNumber string) (models.ServiceProvider, error) {
	var p models.ServiceProvider
	return p, provider.db.Where("phone_number = ?", phoneNumber).First(&p).Error
}

func (provider ServiceProviderAccount) Update(providerId string, data models.ServiceProvider) error {
	c := models.Client{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}
	return provider.db.Where("client_id = ?", providerId).Updates(&c).Error
}

func (provider ServiceProviderAccount) DeactivateAccount(providerId string) error {
	var p models.Client
	return provider.db.Model(&p).Where("client_id = ?", providerId).Delete(p).Error
}

// complete activate account later
func (provider ServiceProviderAccount) ActivateAccount(providerId string) error {
	return nil
}
