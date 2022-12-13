package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientIdentity string
	ClientId       string
	FirstName      string
	LastName       string
	PhoneNumber    string
	Email          string
	Password       string
}

type Address struct {
	gorm.Model
	AddressIdentity string
	ClientId        string
	Name            string
	ZipCode         string
	City            string
}
