package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	SericesIdentity    string
	ServiceProviderId  string
	PhoneNumber        string
	YearOfExperience   string
	Service            string
	CompanyName        string
	Email              string
	IsFullyOnboarded   bool
	IsAccountConfirmed bool
	Pin                string
}

type ServiceJson struct {
	PhoneNumber       string `json:"phone_number" validate:"required"`
	YearOfExperience  string `json:"year_of_experience" validate:"required"`
	Service           string `json:"service" validate:"required"`
	CompanyName       string `json:"company_name" validate:"required"`
	Address           ServiceProviderAddressJson
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

type ServiceProviderDetail struct{
	FullName string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	CompanyName string `json:"company_name"`
	Service string `json:"service"`
	CompanyPhoneNumber string `json:"company_phone_number"`
	AddressName string `json:"address_name"`
	AddressCity string `json:"address_city"`
}
