package handler

import "net/http"

func Services(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "services.html", nil)
}
