//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/patankarcp/ginkgo-poc/pkg/config"
	ghttp "github.com/patankarcp/ginkgo-poc/pkg/http"
	"github.com/patankarcp/ginkgo-poc/pkg/kafka"
	"github.com/patankarcp/ginkgo-poc/pkg/logger"
	"github.com/patankarcp/ginkgo-poc/pkg/server"
)

func InitializeServer() (*server.Server, func()) {
	wire.Build(
		GCommonSet,
		wire.Struct(new(UserService), "*"),
		NewServer,
	)
	return &server.Server{}, nil
}

// This is in a separate common package
var GCommonSet = wire.NewSet(
	NewServerConfig,
	NewServerFactory,
	config.NewAppConfig,
	NewKafkaConfig,
	kafka.NewClient,
	wire.Bind(new(kafka.Logger), new(logger.Logger)),
	logger.NewLogger,
	NewTracer,
	NewHttpServiceConfig,
	ghttp.NewClientProvider,
	wire.Bind(new(ghttp.Logger), new(logger.Logger)),
	ghttp.NewLeveledLogger,
)
