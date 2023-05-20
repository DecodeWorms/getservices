package storage

import (
	"fmt"
	"log"

	"github.com/DecodeWorms/getservices/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	Client *gorm.DB
}

func NewConn(c config.Config) Conn {
	log.Println("establishing database connection")

	uri := fmt.Sprintf("host=%s dbname=%s port=%s password=%s user=%s", c.DatabaseHost, c.DatabaseName, c.DatabasePort, c.DatabasePassword, c.DatabaseUserName)
	panicHandler()
	log.Println("Connecting...")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
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
