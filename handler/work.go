package handler

import "net/http"

func Work(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "work-single.html", nil)
}
