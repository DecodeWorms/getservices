package handler

import (
	"getservices/client"
)

type ServiceProviderHandler struct {
	Provider client.ClientMigrations
}

func NewServiceProviderHandler(p client.ClientMigrations) ServiceProviderHandler {
	return ServiceProviderHandler{
		Provider: p,
	}
}
