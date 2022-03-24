package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home is thi about page handler
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.page.tmpl", map[string]interface{}{
		"name": "Von!",
	})

}

// About is thi about page handler
func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.page.tmpl", map[string]interface{}{
		"name": "Von!",
	})
}
