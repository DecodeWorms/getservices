package server

import (
	"context"
	handler "getservices/handlers"
	"log"
)

type ServiceProviderServer struct {
	Provider handler.ServiceProviderHandler
}

func NewServiceProviderServer(p handler.ServiceProviderHandler) ServiceProviderServer {
	return ServiceProviderServer{
		Provider: p,
	}
}

func (serv ServiceProviderServer) Create(ctx context.Context) {
	if err := serv.Provider.Create(ctx); err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Println(200, "success")

}

func (serv ServiceProviderServer) Address(ctx context.Context) {
	if err := serv.Provider.Address(ctx); err != nil {
		log.Printf("error %v", err)
	}
	log.Println(200, "success")
}
