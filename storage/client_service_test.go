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
	//connect to the virtual PostgreSql on github action
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", "runner")
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

	//create the table client for github action
	if err = cl.Client.AutoMigrate(&models.Client{}); err != nil {
		panic(err)
	}

	client := NewClientAccount(cl.Client)

	data := models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John@mail.com",
	}
	err = client.Create(data)
	assert.NilError(t, err)
}

func TestLogin(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", "runner")
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
	email := "John@mail.com"
	data := models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     email,
	}
	err = client.Create(data)
	assert.NilError(t, err)
	c, err := client.Login(email)
	assert.NilError(t, err)
	assert.Equal(t, c.FirstName, c.FirstName)
}
func TestClients(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", "runner")
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
	_, err = client.Clients()
	assert.NilError(t, err)

}
