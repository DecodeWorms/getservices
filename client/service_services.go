package client

import (
	"getservices/models"
	"time"

	"gorm.io/gorm"
)

type ServiceServices interface {
	Create(cl models.Services) error
	Service(serviceId string) (models.Services, error)
	Services() ([]models.Services, error)
	ServiceByEmail(email string) (models.Services, error)
	ServiceByPhoneNumber(phoneNumber string) (models.Services, error)
	Update(serviceId string, cl models.Client) error
	DeactivateAccount(serviceId string) error
	ActivateAccount(email string) (models.Client, error)
	CreateAddress(add models.ServiceAddress) error
	AddressByServiceId(serviceId string) (models.ServiceAddress, error)
	UpdateAddress(serviceId string) (models.ServiceAddress, error)
}

type ServiceAccount struct {
	db *gorm.DB
}

func NewServiceAccount(db *gorm.DB) ServiceAccount {
	return ServiceAccount{
		db: db,
	}
}

func (service ServiceAccount) Create(serv models.Services) error {
	serv.CreatedAt = time.Now()
	serv.UpdatedAt = time.Now()

	data := models.Services{
		SericesIdentity:   serv.SericesIdentity, // i do not remeber what it does..
		ServiceProviderId: serv.ServiceProviderId,
		PhoneNumber:       serv.PhoneNumber,
		YearOfExperience:  serv.YearOfExperience,
		CompanyName:       serv.CompanyName,
		Email:             serv.Email,
	}
	return service.db.Create(data).Error
}

func (service ServiceAccount) Service(serviceId string) (models.Services, error) {
	var s models.Services
	return s, service.db.Where("client_id = ?", serviceId).First(&s).Error
}

func (service ServiceAccount) Services() ([]models.Services, error) {
	var c []models.Services
	return c, service.db.Find(&c).Error
}

func (service ServiceAccount) ServiceByEmail(email string) (models.Services, error) {
	var s models.Services
	return s, service.db.Where("email = ?", email).First(&s).Error
}

func (service ServiceAccount) ClientByPhoneNumber(phoneNumber string) (models.Services, error) {
	var s models.Services
	return s, service.db.Where("phone_number = ?", phoneNumber).First(&s).Error
}

func (service ServiceAccount) Update(clientId string, data models.Services) error {
	c := models.Services{
		CompanyName:      data.CompanyName,
		YearOfExperience: data.YearOfExperience,
		Email:            data.Email,
		PhoneNumber:      data.PhoneNumber,
	}
	return service.db.Where("client_id = ?", clientId).Updates(&c).Error
}

func (service ServiceAccount) DeactivateAccount(clientId string) error {
	var s models.Client
	return service.db.Model(&s).Where("client_id = ?", clientId).Delete(s).Error
}

// complete activate account later
func (client ServiceAccount) ActivateAccount(clientId string) error {
	return nil
}

func (service ServiceAccount) CreateAddress(add models.ServiceAddress) error {
	ad := models.ServiceAddress{
		ServiceProviderId: add.ServiceProviderId,
		Name:              add.Name,
		ZipCode:           add.ZipCode,
		City:              add.City,
	}
	return service.db.Create(&ad).Error
}

func (service ServiceAccount) AddressByServiceId(serviceId string) (models.ServiceAddress, error) {
	var ad models.ServiceAddress
	return ad, service.db.Where("provider_id = ?", serviceId).First(&ad).Error
}

func (provider ServiceAccount) UpdateAddress(serviceId string, data models.ServiceAddress) error {
	ad := models.ServiceAddress{
		Name:    data.Name,
		ZipCode: data.ZipCode,
		City:    data.City,
	}
	return provider.db.Model(&ad).Where("provider_id = ?", serviceId).Updates(&ad).Error
}
