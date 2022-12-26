package client

import (
	"getservices/models"
	"time"

	"gorm.io/gorm"
)

type ClientServices interface {
	Create(cl models.Client) error
	Client(clientId string) (models.Client, error)
	Clients() ([]models.Client, error)
	ClientByEmail(email string) (models.Client, error)
	ClientByPhoneNumber(phoneNumber string) (models.Client, error)
	Update(clientId string, cl models.Client) error
	DeactivateAccount(clientId string) error
	ActivateAccount(email string) (models.Client, error)
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
	return client.db.Create(data).Error
}

func (client ClientAccount) Client(clientId string) (models.Client, error) {
	var c models.Client
	return c, client.db.Where("client_id = ?", clientId).First(&c).Error
}

func (client ClientAccount) Clients() ([]models.Client, error) {
	var c []models.Client
	return c, client.db.Find(&c).Error
}

func (client ClientAccount) ClientByEmail(email string) (models.Client, error) {
	var c models.Client
	return c, client.db.Where("email = ?", email).First(&c).Error
}

func (client ClientAccount) ClientByPhoneNumber(phoneNumber string) (models.Client, error) {
	var c models.Client
	return c, client.db.Where("phone_number = ?", phoneNumber).First(&c).Error
}

func (client ClientAccount) Update(clientId string, data models.Client) error {
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
	var c models.Client
	return client.db.Model(&c).Where("client_id = ?", clientId).Delete(c).Error
}

// complete activate account later
func (client ClientAccount) ActivateAccount(clientId string) error {
	return nil
}
