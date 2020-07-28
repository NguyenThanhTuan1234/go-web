package service

import "net/http"

type SignUpService interface {
	CreateSignUpPage(http.ResponseWriter, *http.Request) error
}
