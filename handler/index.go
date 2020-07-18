package handler

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	u := GetUser(w, r)
	tpl.ExecuteTemplate(w, "index.html", u)
}
