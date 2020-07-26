package service

import "net/http"

type IndexService interface {
	CreateIndexPage(http.ResponseWriter, *http.Request) error
}
