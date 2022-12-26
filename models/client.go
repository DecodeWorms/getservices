package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientIdentity     string
	ClientId           string
	FirstName          string
	LastName           string
	PhoneNumber        string
	Email              string
	Password           string
	IsFullyOnboarded   bool
	IsAccountConfirmed bool
	Pin                string
}

type ClientJson struct {
	FirstName   string `json:"first_name" gorm:"firstName" validate:"required"`
	LastName    string `json:"last_name" gorm:"lastName" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"phoneNumber" validate:"required"`
	Email       string `json:"email" gorm:"email" validate:"required"`
	Password    string `json:"password" gorm:"password" validate:"required"`
	Pin         string `json:"pin" gorm:"pin" validate:"required"`
	Address     ClientAddressJson
}

type Address struct {
	gorm.Model
	AddressIdentity string
	ClientId        string
	Name            string
	ZipCode         string
	City            string
}

type ClientAddressJson struct {
	Name    string `json:"name" gorm:"name" validate:"required"`
	ZipCode string `json:"zip_code" gorm:"zipCode" validate:"required"`
	City    string `json:"city" gorm:"city" validate:"required"`
}
