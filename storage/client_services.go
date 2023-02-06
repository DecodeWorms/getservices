package storage

import (
	"getservices/models"
	"time"

	"gorm.io/gorm"
)

type ClientServices interface {
	Create(cl models.Client) error
	Login(email string) (*models.Client, error)
	Client(clientId string) (models.Client, error)
	Clients() ([]models.Client, error)
	ClientByEmail(email string) (*models.Client, error)
	ClientByPhoneNumber(phoneNumber string) (*models.Client, error)
	Update(clientId string, cl models.Client) error
	DeactivateAccount(clientId string) error
	ActivateAccount(email string) error
	CreateAddress(add models.Address) error
	AddressByClientId(clientId string) (models.Address, error)
	UpdateAddress(clientId string, data models.Address) error
	GetDeletedAgentById(id string) (*models.Client, error)
	GetDeletdAddressById(id string)(*models.Address, error)
	UpdatePassword(passwordData *models.Client) error
	DeactivateAddress(clientId string)error
	ActivateAddress(clientId string)error
}

type ClientAccount struct {
	db *gorm.DB
}

func NewClientAccount(db *gorm.DB) ClientAccount {
	return ClientAccount{
		db: db,
	}
}

func (client ClientAccount) Create(cl models.Client) error {
	cl.CreatedAt = time.Now()
	cl.UpdatedAt = time.Now()

	data := models.Client{
		ClientIdentity: cl.ClientIdentity, // i do not remeber what it does..
		ClientId:       cl.ClientId,
		FirstName:      cl.FirstName,
		LastName:       cl.LastName,
		PhoneNumber:    cl.PhoneNumber,
		Email:          cl.Email,
		Password:       cl.Password,
	}
	return client.db.Create(&data).Error
}

func (client ClientAccount) Login(email string) (*models.Client, error) {
	var c *models.Client
	return c, client.db.Where("email = ?", email).First(&c).Error

}

func (client ClientAccount) Client(clientId string) (models.Client, error) {
	var c models.Client
	return c, client.db.Where("client_id = ?", clientId).First(&c).Error
}

func (client ClientAccount) Clients() ([]models.Client, error) {
	var c []models.Client
	return c, client.db.Find(&c).Error
}

func (client ClientAccount) ClientByEmail(email string) (*models.Client, error) {
	var c *models.Client
	return c, client.db.Where("email = ?", email).First(&c).Error
}

func (client ClientAccount) ClientByPhoneNumber(phoneNumber string) (*models.Client, error) {
	var c *models.Client
	return c, client.db.Where("phone_number = ?", phoneNumber).First(&c).Error
}

func (client ClientAccount) Update(clientId string, data models.Client) error {
	data.UpdatedAt = time.Now()
	c := models.Client{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}
	return client.db.Where("client_id = ?", clientId).Updates(&c).Error
}

func (client ClientAccount) DeactivateAccount(clientId string) error {
	var c *models.Client
	//c.DeletedAt = time.Now()
	return client.db.Model(&c).Where("client_id = ?", clientId).Delete(&c).Error
}

func(client ClientAccount)DeactivateAddress(clientId string)error{
	var a *models.Address
	return client.db.Model(&a).Where("client_id = ?", clientId).Delete(&a).Error
}

// complete activate account later
func (client ClientAccount) ActivateAccount(clientId string) error {
	return client.db.Model(&models.Client{}).Unscoped().Where("client_id = ?", clientId).Update("deleted_at", nil).Error
}

func(client ClientAccount)ActivateAddress(clientId string)error{
	return client.db.Model(&models.Address{}).Unscoped().Where("client_id = ?",clientId).Update("deleted_at", nil).Error
}

func (client ClientAccount) CreateAddress(add models.Address) error {
	ad := models.Address{
		AddressIdentity: add.AddressIdentity, //probably remove in the future .
		ClientId:        add.ClientId,
		Name:            add.Name,
		ZipCode:         add.ZipCode,
		City:            add.City,
	}
	return client.db.Create(&ad).Error
}

func (client ClientAccount) AddressByClientId(clientId string) (models.Address, error) {
	var ad models.Address
	return ad, client.db.Where("client_id = ?", clientId).First(&ad).Error
}

func (client ClientAccount) UpdateAddress(clientId string, data models.Address) error {
	data.UpdatedAt = time.Now()
	ad := models.Address{
		Name:    data.Name,
		ZipCode: data.ZipCode,
		City:    data.City,
	}
	return client.db.Model(&ad).Where("client_id = ?", clientId).Updates(&ad).Error
}

func (client ClientAccount) GetDeletedAgentById(id string) (*models.Client, error) {
	var c *models.Client
	return c, client.db.Unscoped().Where("client_id = ?", id).First(&c).Error
}

func(client ClientAccount)GetDeletdAddressById(id string)(*models.Address, error){
	var ad *models.Address
	return ad, client.db.Unscoped().Where("client_id = ?",id).First(&ad).Error
}

func (client ClientAccount) UpdatePassword(data *models.Client) error {
	c := &models.Client{
		Password: data.Password,
	}
	return client.db.Model(&c).Where("email = ?", data.Email).Updates(&c).Error
}
