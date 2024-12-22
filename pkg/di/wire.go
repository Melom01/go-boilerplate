//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/Melom01/go-boilerplate/pkg/api"
	"github.com/Melom01/go-boilerplate/pkg/api/handler"
	"github.com/Melom01/go-boilerplate/pkg/config"
	"github.com/Melom01/go-boilerplate/pkg/db"
	"github.com/Melom01/go-boilerplate/pkg/repository"
	"github.com/Melom01/go-boilerplate/pkg/usecase"
)

func InitializeAppDependencies(cfg config.Configuration) (*api.ServerHttp, error) {
	wire.Build(
		api.NewServerHttp,
		db.ConnectDatabase,
		repository.CreateBookRepository,
		usecase.CreateBookUseCase,
		handler.CreateBookHandler,
	)
	return &api.ServerHttp{}, nil
}
