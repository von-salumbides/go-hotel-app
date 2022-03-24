package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/von-salumbides/go-hotel-app/pkg/handlers"
	"github.com/von-salumbides/go-hotel-app/pkg/render"
)

const portNubmer = ":9000"

func main() {
	// Echo instance
	e := echo.New()
	tmpl := render.Templates()
	e.Renderer = tmpl

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", handlers.Home)
	e.GET("/about", handlers.About)

	// Start server
	e.Logger.Fatal(e.Start(portNubmer))
}
