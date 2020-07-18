package handler

import (
	"net/http"
	"time"
)

func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.Un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
