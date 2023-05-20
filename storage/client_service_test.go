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

const user = "runner"

func TestCreate(t *testing.T) {
	//connect to the virtual PostgreSql on github action
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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
		ClientId:    "client-1234-567",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "John@mail.com",
		PhoneNumber: "0900000000000",
	}
	err = client.Create(data)
	assert.NilError(t, err)
}

func TestLogin(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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

func TestClientByEmail(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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
	mockEmail := "John@mail.com"
	mockResponse := &models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John@mail.com",
	}
	res, err := client.ClientByEmail(mockEmail)
	assert.NilError(t, err)
	assert.Equal(t, mockResponse.FirstName, res.FirstName)
}

func TestClientByPhoneNumber(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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

	mockPhoneNumber := "0900000000000"

	cli := &models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John@mail.com",
	}
	res, err := client.ClientByPhoneNumber(mockPhoneNumber)
	assert.NilError(t, err)
	assert.Equal(t, cli.Email, res.Email)
	assert.Equal(t, cli.FirstName, res.FirstName)
}

func TestUpdate(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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
	mockId := "client-1234-567"
	mockUpdatePayload := models.Client{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "John@mail.com",
		PhoneNumber: "0900000000000",
	}

	client := NewClientAccount(cl.Client)
	err = client.Update(mockId, mockUpdatePayload)
	assert.NilError(t, err)

}

func TestDeactivateAccount(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", "abdulhmeed")
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
	mockId := "client-1234-567"

	client := NewClientAccount(cl.Client)
	err = client.DeactivateAccount(mockId)
	assert.NilError(t, err)
}

func TestClient(t *testing.T) {
	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", "localhost", "services", "5432", "password", user)
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
	mockId := "client-1234-567"

	client := NewClientAccount(cl.Client)
	_, err = client.Client(mockId)
	assert.Error(t, err, "record not found")

}
