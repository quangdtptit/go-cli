package httpserver

import (
	"net/http"

	"github.com/quangdtptit/go-cli/config"

	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	App *echo.Echo
	*Server
}

func NewEchoServer(cfg *config.Config, opts ...Option) *EchoServer {
	s := &EchoServer{
		App: nil,
		Server: &Server{
			notify:          make(chan error, 1),
			address:         _defaultAddr,
			readTimeout:     _defaultReadTimeout,
			writeTimeout:    _defaultWriteTimeout,
			idleTimeout:     _defaultIdleTimeout,
			shutdownTimeout: _defaultShutdownTimeout,
		},
	}

	// Apply configuration-based timeouts if provided
	if cfg.HTTP.ReadTimeout > 0 {
		s.readTimeout = cfg.HTTP.ReadTimeout
	}
	if cfg.HTTP.WriteTimeout > 0 {
		s.writeTimeout = cfg.HTTP.WriteTimeout
	}
	if cfg.HTTP.IdleTimeout > 0 {
		s.idleTimeout = cfg.HTTP.IdleTimeout
	}
	if cfg.HTTP.ShutdownTimeout > 0 {
		s.shutdownTimeout = cfg.HTTP.ShutdownTimeout
	}

	// Custom options
	for _, opt := range opts {
		opt(s.Server)
	}

	app := echo.New()
	s.App = app

	// Create HTTP server
	s.srv = &http.Server{
		Addr:         s.address,
		Handler:      s.App,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		IdleTimeout:  s.idleTimeout,
	}

	return s
}
