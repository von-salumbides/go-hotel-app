package render

import (
	"errors"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var functions = template.FuncMap{}

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// RenderTemplate function to render views
func RenderTemplate() *TemplateRegistry {
	templates, err := CreateTemplateCache()
	if err != nil {
		panic(err)
	}
	return &TemplateRegistry{
		templates: templates,
	}
}

// CreateTemplateCache creates the cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	pages, err := filepath.Glob("view/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("view/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("view/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.layout.tmpl", data)
}
