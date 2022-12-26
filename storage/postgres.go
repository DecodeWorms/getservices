package storage

import (
	"fmt"
	"getservices/config"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Conn struct {
	Client *gorm.DB
}

func NewConn(c config.Config) Conn {
	log.Println("establishing database connection")

	uri := fmt.Sprintf("host=%s dbname=%s port=%s", c.DatabaseHost, c.DatabaseName, c.DatabasePort)
	panicHandler()
	log.Println("Connecting...")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Connection established ..")

	conn := Conn{
		Client: db,
	}
	return conn

}

func panicHandler() {

	r := recover()
	if r != nil {
		log.Println("handling panic error :", r)
	}
}
