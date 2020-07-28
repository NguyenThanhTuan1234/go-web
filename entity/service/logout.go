package service

import "net/http"

type LogOutService interface {
	CreateLogOutPage(http.ResponseWriter, *http.Request) error
}
