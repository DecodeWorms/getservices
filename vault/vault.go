package vault

import (
	"getservices/config"
	"os"

	"github.com/joho/godotenv"
)

func GetVault() config.Config {
	_ = godotenv.Load()
	h := os.Getenv("POSTGRES_HOST")
	u := os.Getenv("POSTGRES_USER")
	p := os.Getenv("POSTGRES_PORT")
	n := os.Getenv("POSTGRES_DB")
	pwd := os.Getenv("POSTGRES_PASSWORD")

	c := config.Config{
		DatabaseHost:     h,
		DatabaseUserName: u,
		DatabaseName:     n,
		DatabasePort:     p,
		DatabasePassword: pwd,
	}
	return c
}
