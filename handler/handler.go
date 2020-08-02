package handler

import "net/http"

type Handler interface {
	Index(http.ResponseWriter, *http.Request)
	LogIn(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
	Blog(http.ResponseWriter, *http.Request)
	Contact(http.ResponseWriter, *http.Request)
	Services(http.ResponseWriter, *http.Request)
	Work(http.ResponseWriter, *http.Request)
}
