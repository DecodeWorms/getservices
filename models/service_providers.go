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
}

type ServiceProviderAddress struct {
	gorm.Model
	ServiceProviderAddressIdentity string
	ServiceProviderId              string
	Name                           string
	ZipCode                        string
	City                           string
}
