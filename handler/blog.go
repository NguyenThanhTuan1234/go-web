package handler

import "net/http"

func Blog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "blog.html", nil)
}
