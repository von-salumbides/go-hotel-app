package config

import "text/template"

// TemplateRenderer is a custom html/template renderer for Echo framework
type AppConfig struct {
	templateCache map[string]*template.Template
}
