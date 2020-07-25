package main

import (
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
