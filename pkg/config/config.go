package config

import "html/template"

// AppConffig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
