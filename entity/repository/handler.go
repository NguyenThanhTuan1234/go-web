package repository

import "net/http"

type HandlerRepository interface {
	Index(http.ResponseWriter, *http.Request) error
	LogIn(http.ResponseWriter, *http.Request) error
	SignUp(http.ResponseWriter, *http.Request) error
	Blog(http.ResponseWriter, *http.Request) error
}
