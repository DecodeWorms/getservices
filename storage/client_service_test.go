package storage

import (
	"fmt"
	"log"
	"testing"

	"github.com/DecodeWorms/getservices/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "", "runner")
	log.Println("Connecting...")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Connection established ..")

	cl := Conn{
		Client: db,
	}

	client := NewClientAccount(cl.Client)

	data := models.Client{
		FirstName: "John",
		LastName:  "Doe",
	}
	err = client.Create(data)
	assert.NilError(t, err)
}

func TestLogin(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "", "runner")
	log.Println("Connecting...")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Connection established ..")

	cl := Conn{
		Client: db,
	}

	client := NewClientAccount(cl.Client)
	email := "john@mail.com"
	data := models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     email,
	}
	err = client.Create(data)
	assert.NilError(t, err)
	c, err := client.Login(email)
	assert.NilError(t, err)
	assert.Equal(t, c.FirstName, data.FirstName)
}
