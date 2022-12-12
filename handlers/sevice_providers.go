package handler

import (
	"context"
	"fmt"
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

func (serv ServiceProviderHandler) Create(ctx context.Context) error {
	if err := serv.Provider.Create(ctx); err != nil {
		return fmt.Errorf("error migrating service provider table %v", err)
	}
	return nil
}

func (serv ServiceProviderHandler) Address(ctx context.Context) error {
	if err := serv.Provider.Address(ctx); err != nil {
		return fmt.Errorf("error migrating service provider table %v", err)
	}
	return nil
}
