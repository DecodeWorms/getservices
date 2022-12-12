package main

import (
	"os"

	"getservices/config"
	"getservices/storage"

	"github.com/joho/godotenv"
)

var db storage.Conn

func init() {
	_ = godotenv.Load()
	h := os.Getenv("DATABASE_HOST")
	u := os.Getenv("DATABASE_USERNAME")
	p := os.Getenv("DATABASE_PORT")
	n := os.Getenv("DATABASE_NAME")

	c := config.Config{
		DatabaseHost:     h,
		DatabaseUserName: u,
		DatabaseName:     n,
		DatabasePort:     p,
	}
	db = storage.NewConn(c)

}

func main() {

}
