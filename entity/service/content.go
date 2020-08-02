package service

import "net/http"

type ContentService interface {
	CreateContentPage(http.ResponseWriter, *http.Request) error
}
