package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/quangdtptit/go-cli/config"
	"github.com/quangdtptit/go-cli/internal/app/middleware"
	"github.com/quangdtptit/go-cli/internal/modules/user"
	"github.com/quangdtptit/go-cli/pkg/httpserver"
	"github.com/quangdtptit/go-cli/pkg/logger"
	"github.com/quangdtptit/go-cli/pkg/pprof"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	// http server
	echoHttpServer := httpserver.NewEchoServer(cfg, httpserver.Port(cfg.HTTP.Port))
	echoHttpServer.Start()
	logger.Logger.Info(fmt.Sprintf("app running at port http:%s", cfg.HTTP.Port))

	// http server > middleware
	echoHttpServer.App.Use(middleware.LoggerMiddleware())

	// controller
	apiV1 := echoHttpServer.App.Group("/api/v1")
	userController := user.NewController(user.NewService())
	userController.Register(apiV1)

	// pprof
	pprof.Start(cfg)

	// system
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		logger.Logger.Info(fmt.Sprintf("app - Run - signal: %s", s.String()))
	case err := <-echoHttpServer.Notify():
		logger.Logger.Error("app - Run - EchoHttpServer.Notify", zap.Error(err))
	}
	err := echoHttpServer.Shutdown()
	if err != nil {
		logger.Logger.Error("app - Run - EchoHttpServer.Shutdown: ", zap.Error(err))
	}
}
