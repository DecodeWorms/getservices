package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	SericesIdentity   string
	ServiceProviderId string
	PhoneNumber       string
	YearOfExperience  string
	CompanyName       string
	Email             string
}

type ServiceAddress struct {
	gorm.Model
	ServicesAddressIdentity string
	ServiceProviderId       string
	Name                    string
	City                    string
	ZipCode                 string
}
