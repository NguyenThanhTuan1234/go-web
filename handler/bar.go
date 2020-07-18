package handler

import "net/http"

func Bar(w http.ResponseWriter, r *http.Request) {
	u := GetUser(w, r)
	if !AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.html", u)
}
