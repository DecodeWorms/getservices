package models

import "gorm.io/gorm"

type ServiceProvider struct {
	gorm.Model
	ClientId    string `json:"client_id" gorm:"client_id"`
	FirstName   string `json:"first_name" gorm:"first_name"`
	LastName    string `json:"last_name" gorm:"first_name"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	Email       string `json:"email" gorm:"email"`
	Password    string `json:"password" gorm:"password"`
}

type ServiceProviderAddress struct {
	gorm.Model
	ClientId string `json:"client_id" gorm:"client_id"`
	Name     string `json:"name" gorm:"name"`
	ZipCode  string `json:"zip_code" gorm:"zip_code"`
	City     string `json:"city" gorm:"city"`
}
