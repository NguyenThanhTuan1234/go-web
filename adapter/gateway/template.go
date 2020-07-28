package gateway

import (
	"go-web/entity/repository"
	"html/template"
)

type templateClient struct {
	tpl *template.Template
}

type TemplateRepository interface {
	repository.HandlerRepository
}

func NewTemplateRepository() TemplateRepository {
	tpl := template.Must(template.ParseGlob("public/*"))
	return &templateClient{
		tpl: tpl,
	}
}
