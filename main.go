package main

import (
	"context"
	"log"
	"os"

	"getservices/client"
	"getservices/config"
	handler "getservices/handlers"
	"getservices/server"
	"getservices/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db storage.Conn
var clientStorage client.Client
var serviceProviderStorage client.ServiceProvider
var serviceStorage client.Service

var clientServer server.ClientMigrationServer
var serviceProviderServer server.ServiceProviderServer
var serviceServer server.ServicesServer

var clientHandler handler.ClientMigrationHandler
var serviceProviderHandler handler.ServiceProviderHandler
var serviceHandler handler.ServicesHandler

func init() {

	var ctx context.Context
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
	clientStorage = client.NewCleint(db.Client)
	serviceProviderStorage = client.NewServeProvider(db.Client)
	serviceStorage = client.NewServices(db.Client)

	clientHandler = handler.NewCleintMigration(clientStorage)
	serviceProviderHandler = handler.NewServiceProviderHandler(serviceProviderStorage)
	serviceHandler = handler.NewServiceHandler(serviceStorage)

	clientServer = server.NewClientMigrationServer(clientHandler)
	serviceProviderServer = server.NewServiceProviderServer(serviceProviderHandler)
	serviceServer = server.NewServiceServer(serviceHandler)

	clientServer.Create(ctx)
	clientServer.Address(ctx)

	serviceProviderServer.Create(ctx)
	serviceProviderServer.Address(ctx)

	serviceServer.Create(ctx)
	serviceServer.Address(ctx)

}

func main() {
	router := gin.New()

	if err := router.Run(":080"); err != nil {
		log.Println("error processing http server req")
	}

}
