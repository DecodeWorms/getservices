package server

import (
	handler "getservices/handlers"
)

type ServicesServer struct {
	ServiceServer handler.ServicesHandler
}

func NewServiceServer(s handler.ServicesHandler) ServicesServer {
	return ServicesServer{
		ServiceServer: s,
	}
}
