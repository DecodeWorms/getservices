package server

import (
	"context"
	"log"

	handler "github.com/DecodeWorms/getservices/handlers"
)

type ClientMigrationServer struct {
	ClientHandler handler.ClientMigrationHandler
}

func NewClientMigrationServer(cl handler.ClientMigrationHandler) ClientMigrationServer {
	return ClientMigrationServer{
		ClientHandler: cl,
	}

}

func (cl ClientMigrationServer) MigrateModels(ctx context.Context) {

	if err := cl.ClientHandler.MigrateModels(ctx); err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Println(200, "success")

}
