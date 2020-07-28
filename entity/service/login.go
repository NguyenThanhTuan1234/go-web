package service

import "net/http"

type LoginService interface {
	CreateLoginPage(http.ResponseWriter, *http.Request) error
}
