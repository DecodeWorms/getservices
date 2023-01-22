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
var providerService storage.ProviderServices
var serviceService storage.ServiceAccount

// handler for migration tables
var clientHandler handler.ClientMigrationHandler

// server handler for table migrations
var clientServer server.ClientMigrationServer

// server handler for clients , providers and services
var clientServ server.ClientServer
var providerServer server.ProviderServer
var serviceServer server.ServiceServer

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
	providerService = storage.NewServiceProviderAccount(db.Client)
	serviceService = storage.NewServiceAccount(db.Client)

	clientHand := handler.NewClientHandler(clientService)
	providerHandler := handler.NewServiceProviderHandler(providerService)
	serviceHandler := handler.NewServiceHandler(serviceService,providerService)

	clientServ = server.NewClientServer(clientHand)
	providerServer = server.NewProviderServer(providerHandler)
	serviceServer = server.NewServiceServer(serviceHandler)

}

func main() {
	router := gin.New()

	//client public api endpoints
	router.POST("/client", clientServ.SignUpClient())
	router.POST("/client/login", clientServ.UserLogin())
	router.PUT("/client", clientServ.UpdateClient())
	router.DELETE("/client", clientServ.DeactivateAccount())
	router.PUT("/client/reactivate", clientServ.ActivateAccount())
	router.PUT("client/update_password", clientServ.UpdateClientPassword())

	//provider public api endpoint
	router.POST("/provider", providerServer.SignUpProvider())
	router.POST("/provider/login", providerServer.LoginProvider())
	router.PUT("provider/update_password", providerServer.UpdatePassword())

	//service public endpoint
	router.GET("/service/categories",serviceServer.GetServicesCategories())
	router.GET("/services",serviceServer.GetServices())
	router.POST("/service", serviceServer.CreateService())

	if err := router.Run(":080"); err != nil {
		log.Println("error processing http server req")
	}

}
