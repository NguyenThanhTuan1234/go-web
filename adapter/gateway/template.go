package gateway

import (
	"go-web/entity/repository"
	"html/template"
)

type templateClient struct {
	tpl *template.Template
}

type TemplateClient interface {
	repository.HandlerRepository
}

func NewTemplateClient() TemplateClient {
	tpl := template.Must(template.ParseGlob("public/*"))
	return &templateClient{
		tpl: tpl,
	}
}
