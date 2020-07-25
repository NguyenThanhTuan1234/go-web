package service

import "net/http"

type LoginService interface {
	CreateLoginPage(w http.ResponseWriter, r *http.Request) error
}
