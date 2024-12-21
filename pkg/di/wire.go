//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Melom01/go-boilerplate/pkg/api"
	"github.com/Melom01/go-boilerplate/pkg/config"

	"github.com/google/wire"
)

func InitializeAppDependencies(cfg config.Configuration) (*api.ServerHttp, error) {
	wire.Build(
		api.NewServerHttp,
	)
	return &api.ServerHttp{}, nil
}
