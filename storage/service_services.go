package storage

import (
	"time"

	"github.com/DecodeWorms/getservices/models"

	"gorm.io/gorm"
)

type ServiceServices interface {
	Create(cl models.Services) error
	Service(serviceId string) (*models.Services, error)
	Services() ([]models.Services, error)
	ServiceByEmail(email string) (models.Services, error)
	ServiceByCompanyName(companyName string) (*models.Services, error)
	ServiceByService(serviceName string) ([]*models.Services, error)
	ServiceByPhoneNumber(phoneNumber string) (*models.Services, error)
	Update(serviceId string, cl models.Services) error
	DeactivateAccount(serviceId string) error
	ActivateAccount(email string) error
	CreateAddress(add models.ServiceAddress) error
	AddressByProviderId(serviceId string) (models.ServiceAddress, error)
	UpdateAddress(serviceProviderId string, data *models.ServiceAddress) error
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
		SericesIdentity:   serv.SericesIdentity, // i do not remember what it does..
		ServiceProviderId: serv.ServiceProviderId,
		PhoneNumber:       serv.PhoneNumber,
		YearOfExperience:  serv.YearOfExperience,
		Service:           serv.Service,
		CompanyName:       serv.CompanyName,
	}
	return service.db.Create(&data).Error
}

func (service ServiceAccount) Service(serviceId string) (*models.Services, error) {
	var s *models.Services
	return s, service.db.Where("service_provider_id = ?", serviceId).First(&s).Error
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

func (service ServiceAccount) Update(serviceProviderId string, data models.Services) error {
	data.UpdatedAt = time.Now()
	c := models.Services{
		CompanyName:      data.CompanyName,
		YearOfExperience: data.YearOfExperience,
		Email:            data.Email,
		PhoneNumber:      data.PhoneNumber,
	}
	return service.db.Where("service_provider_id = ?", serviceProviderId).Updates(&c).Error
}

func (service ServiceAccount) DeactivateAccount(clientId string) error {
	var s models.Client
	return service.db.Model(&s).Where("client_id = ?", clientId).Delete(s).Error
}

// complete activate account later
func (service ServiceAccount) ActivateAccount(clientId string) error {
	return nil
}

func (service ServiceAccount) ServiceByPhoneNumber(phoneNumber string) (*models.Services, error) {
	var s *models.Services
	return s, service.db.Where("phone_number = ?", phoneNumber).First(&s).Error
}

func (service ServiceAccount) ServiceByService(serviceCat string) ([]*models.Services, error) {
	var s []*models.Services
	return s, service.db.Where("service = ?", serviceCat).Find(&s).Error
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

func (service ServiceAccount) AddressByProviderId(providerId string) (models.ServiceAddress, error) {
	var ad models.ServiceAddress
	return ad, service.db.Where("service_provider_id = ?", providerId).First(&ad).Error
}

func (service ServiceAccount) UpdateAddress(serviceProviderId string, data *models.ServiceAddress) error {
	data.UpdatedAt = time.Now()
	ad := &models.ServiceAddress{
		Name:    data.Name,
		ZipCode: data.ZipCode,
		City:    data.City,
	}
	return service.db.Model(&ad).Where("service_provider_id = ?", serviceProviderId).Updates(&ad).Error
}

func (service ServiceAccount) ServiceByCompanyName(companyName string) (*models.Services, error) {
	var s *models.Services
	return s, service.db.Where("company_name = ?", companyName).First(&s).Error
}
