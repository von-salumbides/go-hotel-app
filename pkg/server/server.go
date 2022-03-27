package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/von-salumbides/go-hotel/pkg/render"
	"go.uber.org/zap"
)

type ServerConfig struct {
	httpAddress string
	debug       bool
}

type Server struct {
	httpAddress string
	Echo        *echo.Echo
}

func NewServerConfig(httpAddress string, debug bool) *ServerConfig {
	return &ServerConfig{
		httpAddress: httpAddress,
		debug:       debug,
	}
}

func NewServer(config *ServerConfig) (*Server, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Renderer = render.RenderTemplate()
	// Log HTTP requests
	e.Use(middleware.Logger())
	return &Server{
		httpAddress: config.httpAddress,
		Echo:        e,
	}, nil
}

func (server *Server) Start() {
	zap.L().Info("Starting HTTP server",
		zap.String("address", server.httpAddress),
	)

	err := server.Echo.Start(server.httpAddress)
	if err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Failed to start server",
			zap.String("address", server.httpAddress),
			zap.Error(err))
	}
}

func (server *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	logger := zap.L()
	logger.Info("Shutting down HTTP server",
		zap.String("address", server.httpAddress))
	if err := server.Echo.Shutdown(ctx); err != nil {
		logger.Error("Failed to shut down HTTP server gracefully", zap.Error(err))
		return err
	}
	return nil
}
