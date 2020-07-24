package repository

import "net/http"

type HandlerRepository interface {
	Index(http.ResponseWriter, *http.Request)
	LogIn(http.ResponseWriter, *http.Request)
}
