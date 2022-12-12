package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	UserId      string `json:"user_id" gorm:"user_id"`
	FirstName   string `json:"first_name" gorm:"first_name"`
	LastName    string `json:"last_name" gorm:"first_name"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	Email       string `json:"email" gorm:"email"`
	Password    string `json:"password" gorm:"password"`
}

type Address struct {
	gorm.Model
	UserId  string `json:"user_id" gorm:"user_id"`
	Name    string `json:"name" gorm:"name"`
	ZipCode string `json:"zip_code" gorm:"zip_code"`
	City    string `json:"city" gorm:"city"`
}
