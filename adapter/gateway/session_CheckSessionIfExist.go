package gateway

import (
	"net/http"
	"time"
)

func (e *sessionRepository) CheckSessionIfExist(w http.ResponseWriter, r *http.Request, name string) bool {
	c, err := r.Cookie(name)
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
