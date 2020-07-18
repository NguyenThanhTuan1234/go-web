package handler

import (
	"go-web/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetUser(w http.ResponseWriter, r *http.Request) models.User {

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	var u models.User
	// if the user exists already, get user
	if s, ok := dbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.Un]
	}
	return u
}
