package models

type TemplateData struct {
	TemplateData map[string]interface{}
	CSRFToken    string
	Flash        string
	Warning      string
	Error        string
}
