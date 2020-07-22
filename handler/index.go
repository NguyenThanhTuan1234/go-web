package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	u := GetUser(w, r)
	fmt.Println(u)
	tpl.ExecuteTemplate(w, "index.html", u)
}
