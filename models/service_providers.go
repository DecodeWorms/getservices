package models

import "gorm.io/gorm"

type ServiceProvider struct {
	gorm.Model
	ServiceProviderIdentity string
	ServiceProviderId       string
	FirstName               string
	LastName                string
	PhoneNumber             string
	Email                   string
	Password                string
	IsFullyOnboarded        bool
	IsAccountConfirmed      bool
	Pin                     string
}

type ServiceProviderJson struct {
	FirstName       string `json:"first_name" gorm:"firstName" validate:"required"`
	LastName        string `json:"last_name" gorm:"lastName" validate:"required"`
	PhoneNumber     string `json:"phone_number" gorm:"phoneNumber" validate:"required"`
	Email           string `json:"email" gorm:"email" validate:"required"`
	Password        string `json:"password" gorm:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" gorm:"confirm_password" validate:"required"`
	//Pin             string                     `json:"pin" gorm:"pin" validate:"required"`
	Address ServiceProviderAddressJson `json:"address" gorm:"address" validate:"required"`
}

type ServiceProviderAddress struct {
	gorm.Model
	ServiceProviderAddressIdentity string
	ServiceProviderId              string
	Name                           string
	ZipCode                        string
	City                           string
}

type ServiceProviderAddressJson struct {
	Name    string `json:"name" gorm:"name" validate:"required"`
	ZipCode string `json:"zip_code" gorm:"zipCode" validate:"required"`
	City    string `json:"city" gorm:"city" validate:"required"`
}
