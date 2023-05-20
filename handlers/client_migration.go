package handlers

import (
	"context"
	"fmt"

	"github.com/DecodeWorms/getservices/migrator"

	"golang.org/x/sync/errgroup"
)

type ClientMigrationHandler struct {
	ClientMig        migrator.ClientMigrations
	ServiceMig       migrator.ServiceMigrations
	ServiceProvidmig migrator.ServiceProviderMigrations
}

func NewCleintMigration(cl migrator.ClientMigrations, srvMig migrator.ServiceMigrations, srvProMig migrator.ServiceProviderMigrations) ClientMigrationHandler {
	return ClientMigrationHandler{
		ClientMig:        cl,
		ServiceMig:       srvMig,
		ServiceProvidmig: srvProMig,
	}

}

func (cl ClientMigrationHandler) MigrateModels(ctx context.Context) error {

	var g errgroup.Group
	//client tables
	g.Go(func() error {
		if err := cl.ClientMig.Create(ctx); err != nil {
			return fmt.Errorf("error generating client table %v", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := cl.ClientMig.Address(ctx); err != nil {
			return fmt.Errorf("error generating client address table %v", err)

		}
		return nil
	})

	//services tables

	g.Go(func() error {
		if err := cl.ServiceMig.Create(ctx); err != nil {
			return fmt.Errorf("error generating client table %v", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := cl.ServiceMig.Address(ctx); err != nil {
			return fmt.Errorf("error generating client address table %v", err)
		}
		return nil
	})

	//service provider tables
	g.Go(func() error {
		if err := cl.ServiceProvidmig.Create(ctx); err != nil {
			return fmt.Errorf("error generating client address table %v", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := cl.ServiceProvidmig.Address(ctx); err != nil {
			return fmt.Errorf("error generating client address table %v", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return err
	}

	return nil

}
