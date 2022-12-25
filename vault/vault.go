package vault

import (
	"getservices/config"
	"os"

	"github.com/joho/godotenv"
)

func GetVault() config.Config {
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
	return c
}
