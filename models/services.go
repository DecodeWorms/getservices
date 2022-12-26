package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	SericesIdentity    string
	ServiceProviderId  string
	PhoneNumber        string
	YearOfExperience   string
	CompanyName        string
	Email              string
	IsFullyOnboarded   bool
	IsAccountConfirmed bool
	Pin                string
}

type ServiceJson struct {
	FirstName   string `json:"first_name" gorm:"firstName" validate:"required"`
	LastName    string `json:"last_name" gorm:"lastName" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"phoneNumber" validate:"required"`
	Email       string `json:"email" gorm:"email" validate:"required"`
	Password    string `json:"password" gorm:"password" validate:"required"`
	Pin         string `json:"pin" gorm:"pin" validate:"required"`
	Address     ServiceProviderAddressJson
}

type ServiceAddress struct {
	gorm.Model
	ServicesAddressIdentity string
	ServiceProviderId       string
	Name                    string
	City                    string
	ZipCode                 string
}

type ServiceAddressJson struct {
	Name    string `json:"name" gorm:"name" validate:"required"`
	ZipCode string `json:"zip_code" gorm:"zipCode" validate:"required"`
	City    string `json:"city" gorm:"city" validate:"required"`
}
