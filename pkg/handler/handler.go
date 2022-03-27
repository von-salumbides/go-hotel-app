package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/von-salumbides/go-hotel/pkg/models"
)

func CreateTemplateData(td *models.TemplateData) interface{} {
	return td
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.page.tmpl", map[string]interface{}{
		"name": "Home",
		"msg":  "This is the Home Page",
	})
}
func AboutHandler(c echo.Context) error {
	var data models.TemplateData
	data.TemplateData = map[string]interface{}{
		"name": "About",
		"msg":  "This is the About Page",
	}
	CreateTemplateData(&data)
	return c.Render(http.StatusOK, "about.page.tmpl", data.TemplateData)
}
