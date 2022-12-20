package server

import (
	handler "getservices/handlers"
)

type ServiceProviderServer struct {
	Provider handler.ServiceProviderHandler
}

func NewServiceProviderServer(p handler.ServiceProviderHandler) ServiceProviderServer {
	return ServiceProviderServer{
		Provider: p,
	}
}
