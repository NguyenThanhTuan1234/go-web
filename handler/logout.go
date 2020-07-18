package handler

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if !AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	c, _ := r.Cookie("session")

	// delete the session
	delete(dbSessions, c.Value)

	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Minute * 10) {
		go CleanSessions()
	}

	http.Redirect(w, r, "login", http.StatusSeeOther)
}
