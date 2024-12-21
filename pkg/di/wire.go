//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/Melom01/go-boilerplate/pkg/api"
	"github.com/Melom01/go-boilerplate/pkg/config"
)

func InitializeAppDependencies(cfg config.Configuration) (*api.ServerHttp, error) {
	wire.Build(
		api.NewServerHttp,
	)
	return &api.ServerHttp{}, nil
}
