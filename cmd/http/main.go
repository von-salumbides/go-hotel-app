package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/von-salumbides/go-hotel/pkg/router"
	"github.com/von-salumbides/go-hotel/pkg/server"
	"go.uber.org/zap"
)

func main() {
	logger := zap.L()
	// Setup
	// e := echo.New()
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	// e.Renderer = render.RenderTemplate()

	// Route => handler
	// e.GET("/about", handler.AboutHandler)
	// e.GET("/", handler.HomeHandler)

	// Start server
	// go func() {
	// 	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
	// 		e.Logger.Fatal("shutting down the server")
	// 	}
	// }()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := e.Shutdown(ctx); err != nil {
	// 	e.Logger.Fatal(err)
	// }
	serverConfig := server.NewServerConfig(":8080", false)
	httpServer, err := server.NewServer(serverConfig)
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("Shutting down the server", zap.Error(err))
		os.Exit(1)
	}
	router.Routes(httpServer)
	go httpServer.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	err = httpServer.Shutdown(300 * time.Second)
	if err == nil {
		logger.Info("Shutdown successfully")
	}
}
