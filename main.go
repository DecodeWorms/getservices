package main

import (
	"context"
	"log"

	"getservices/client"
	handler "getservices/handlers"
	"getservices/server"
	"getservices/storage"
	"getservices/vault"

	"github.com/gin-gonic/gin"
)

var db storage.Conn
var clientStorage client.Client
var serviceProviderStorage client.ServiceProvider
var serviceStorage client.Service

var clientHandler handler.ClientMigrationHandler
var serviceProviderHandler handler.ServiceProviderHandler
var serviceHandler handler.ServicesHandler

var clientServer server.ClientMigrationServer
var serviceProviderServer server.ServiceProviderServer
var serviceServer server.ServicesServer

func init() {
	var ctx context.Context
	c := vault.GetVault()
	db = storage.NewConn(c)
	clientStorage = client.NewCleint(db.Client)
	serviceProviderStorage = client.NewServeProvider(db.Client)
	serviceStorage = client.NewServices(db.Client)

	clientHandler = handler.NewCleintMigration(clientStorage, serviceStorage, serviceProviderStorage)
	serviceProviderHandler = handler.NewServiceProviderHandler(serviceProviderStorage)
	serviceHandler = handler.NewServiceHandler(serviceStorage)

	clientServer = server.NewClientMigrationServer(clientHandler)
	serviceProviderServer = server.NewServiceProviderServer(serviceProviderHandler)
	serviceServer = server.NewServiceServer(serviceHandler)
	clientServer.MigrateModels(ctx)

}

func main() {
	router := gin.New()

	if err := router.Run(":080"); err != nil {
		log.Println("error processing http server req")
	}

}
