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
	Password           string `gorm:"type:varchar(70)"`
	IsFullyOnboarded   bool
	IsAccountConfirmed bool
	Pin                string
}

type ClientJson struct {
	FirstName       string `json:"first_name" gorm:"first_name" validate:"required"`
	LastName        string `json:"last_name" gorm:"last_name" validate:"required"`
	PhoneNumber     string `json:"phone_number" gorm:"phone_number" validate:"required"`
	Email           string `json:"email" gorm:"email" validate:"required"`
	Password        string `json:"password" gorm:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" gorm:"confirm_password" validate:"required"`
	//Pin         string `json:"pin" gorm:"pin" validate:"required"`
	Address ClientAddressJson `json:"address"`
}

type ClientLoginJson struct {
	Email    string `json:"email" gorm:"email" validate:"required"`
	Password string `json:"password" gorm:"password" validate:"required"`
}

type ClientLoginResponse struct {
	ClientIdentity     string
	ClientId           string            `json:"client_id"`
	FirstName          string            `json:"first_name"`
	LastName           string            `json:"last_name"`
	PhoneNumber        string            `json:"phone_number"`
	Email              string            `json:"email" gorm:"email" validate:"required"`
	IsFullyOnboarded   bool              `json:"is_fullyonboarded"`
	IsAccountConfirmed bool              `json:"is_account_confirmed"`
	Pin                string            `json:"pin"`
	Address            ClientAddressJson `json:"address"`
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
	ZipCode string `json:"zip_code" gorm:"zip_code" validate:"required"`
	City    string `json:"city" gorm:"city" validate:"required"`
}

type PasswordJson struct {
	Password           string `json:"password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}
