package server

import (
	"context"
	handler "getservices/handlers"
	"log"
)

type ServicesServer struct {
	ServiceServer handler.ServicesHandler
}

func NewServiceServer(s handler.ServicesHandler) ServicesServer {
	return ServicesServer{
		ServiceServer: s,
	}
}

func (s ServicesServer) Create(ctx context.Context) {
	if err := s.ServiceServer.Create(ctx); err != nil {
		log.Printf("error :%v", err)
		return
	}
	log.Println(200, "success")
}

func (s ServicesServer) Address(ctx context.Context) {
	if err := s.ServiceServer.Address(ctx); err != nil {
		log.Printf("error %v:", err)
		return
	}
	log.Println(200, "success")

}
