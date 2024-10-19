//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/spajam-backend/db"
	"github.com/K-Kizuku/spajam-backend/internal/app/handler"
	"github.com/K-Kizuku/spajam-backend/internal/app/repository"
	"github.com/K-Kizuku/spajam-backend/internal/app/service"
	"github.com/K-Kizuku/spajam-backend/pkg/gcp"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		db.New,
		gcp.NewStorageClient,
		repository.NewUserRepository,
		repository.NewStorageRepository,
		service.NewUserService,
		handler.NewUserHandler,
		handler.New,
	)
	return &handler.Root{}
}
