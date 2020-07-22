package main

import (
	"go-web/handler"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/bar", handler.Bar)
	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/blog", handler.Blog)
	http.HandleFunc("/work", handler.Work)
	http.HandleFunc("/contact", handler.Contact)
	http.HandleFunc("/services", handler.Services)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
