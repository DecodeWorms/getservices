package handler

import (
	"context"
	"fmt"
	"getservices/client"
)

type ClientMigrationHandler struct {
	ClientMig client.ClientMigrations
}

func NewCleintMigration(cl client.ClientMigrations) ClientMigrationHandler {
	return ClientMigrationHandler{
		ClientMig: cl,
	}

}

func (cl ClientMigrationHandler) Create(ctx context.Context) error {
	err := cl.ClientMig.Create(ctx)
	if err != nil {
		return fmt.Errorf("error generating client table %v", err)
	}
	return nil

}

func (cl ClientMigrationHandler) Address(ctx context.Context) error {
	if err := cl.ClientMig.Address(ctx); err != nil {
		return fmt.Errorf("error generating client address table %v", err)
	}
	return nil
}
