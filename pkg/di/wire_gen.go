// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Melom01/go-boilerplate/pkg/api"
	"github.com/Melom01/go-boilerplate/pkg/config"
)

// Injectors from wire.go:

func InitializeAppDependencies(cfg config.Configuration) (*api.ServerHttp, error) {
	serverHttp := api.NewServerHttp()
	return serverHttp, nil
}
