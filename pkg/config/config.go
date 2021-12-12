package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCacge      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
