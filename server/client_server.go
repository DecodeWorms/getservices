package server

import (
	"context"
	handler "getservices/handlers"
	"log"
)

type ClientMigrationServer struct {
	ClientHandler handler.ClientMigrationHandler
}

func NewClientMigrationServer(cl handler.ClientMigrationHandler) ClientMigrationServer {
	return ClientMigrationServer{
		ClientHandler: cl,
	}

}

func (cl ClientMigrationServer) Create(ctx context.Context) {

	err := cl.ClientHandler.Create(ctx)
	if err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Println(200, "success")

}

func (cl ClientMigrationServer) Address(ctx context.Context) {

	if err := cl.ClientHandler.Address(ctx); err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Println(200, "success")
}
