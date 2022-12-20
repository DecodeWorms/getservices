package handler

import (
	"getservices/client"
)

type ServicesHandler struct {
	Services client.ClientMigrations
}

func NewServiceHandler(s client.ClientMigrations) ServicesHandler {
	return ServicesHandler{
		Services: s,
	}
}
