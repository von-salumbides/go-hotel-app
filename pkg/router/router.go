package router

import (
	"github.com/von-salumbides/go-hotel/pkg/handler"
	"github.com/von-salumbides/go-hotel/pkg/server"
)

func Routes(server *server.Server) {
	r := server.Echo

	r.GET("/", handler.HomeHandler)
	r.GET("/about", handler.AboutHandler)

}
