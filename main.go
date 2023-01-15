package main

import (
	"context"
	"log"

	handler "getservices/handlers"
	"getservices/migrator"
	"getservices/server"
	"getservices/storage"
	"getservices/vault"

	"github.com/gin-gonic/gin"
)

var db storage.Conn

// services for database migrations
var clientStorage migrator.Client
var serviceProviderStorage migrator.ServiceProvider
var serviceStorage migrator.Service

// services for clients and providers
var clientService storage.ClientAccount
var providerServ storage.ProviderServices

// handler for migration tables
var clientHandler handler.ClientMigrationHandler

// handler for client, service and service providers
var clientHand handler.ClientHandler

// server handler for table migrations
var clientServer server.ClientMigrationServer

// server handler for clients , providers and services
var clientServ server.ClientServer

func init() {
	var ctx context.Context
	c := vault.GetVault()
	db = storage.NewConn(c)

	//handling table migration
	clientStorage = migrator.NewCleint(db.Client)
	serviceStorage = migrator.NewServices(db.Client)
	serviceProviderStorage = migrator.NewServeProvider(db.Client)
	clientHandler = handler.NewCleintMigration(clientStorage, serviceStorage, serviceProviderStorage)
	clientServer = server.NewClientMigrationServer(clientHandler)
	clientServer.MigrateModels(ctx)

	clientService = storage.NewClientAccount(db.Client)
	//serviceProvid := client.NewServiceProviderAccount(db.Client)

	clientHand := handler.NewClientHandler(clientService)

	clientServ = server.NewClientServer(clientHand)

}

func main() {
	router := gin.New()
	router.POST("/client", clientServ.SignUpClient())
	router.POST("/client/login", clientServ.UserLogin())
	router.PUT("/client", clientServ.UpdateClient())
	router.DELETE("/client", clientServ.DeactivateAccount())
	router.PUT("/client/reactivate", clientServ.ActivateAccount())
	router.PUT("client/update_password", clientServ.UpdateClientPassword())

	if err := router.Run(":080"); err != nil {
		log.Println("error processing http server req")
	}

}
