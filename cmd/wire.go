//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mvp/config"
	"mvp/internal/data"
	"mvp/internal/handler"
	"mvp/internal/server"
	"mvp/utils/app"
	"mvp/utils/log"

	"github.com/google/wire"
)

// initApp 初始化应用
func initApp(log log.Logger, conf *config.Config) (*app.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		handler.ProviderSet,
		newApp,
	))
}
