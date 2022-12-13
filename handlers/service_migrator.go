package handler

import (
	"context"
	"fmt"
	"getservices/client"
)

type ServicesHandler struct {
	Services client.Service
}

func NewServiceHandler(s client.Service) ServicesHandler {
	return ServicesHandler{
		Services: s,
	}
}

func (s ServicesHandler) Create(ctx context.Context) error {
	if err := s.Services.Create(ctx); err != nil {
		return fmt.Errorf("error migrating table for services %v", err)
	}
	return nil
}

func (s ServicesHandler) Address(ctx context.Context) error {
	if err := s.Services.Address(ctx); err != nil {
		return fmt.Errorf("error migrating table for services address %v", err)
	}
	return nil
}
