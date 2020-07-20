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
		rows, err := db.Query("SELECT username, password, first, last, role FROM test1 WHERE username = $1", s.Un)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		}
		for rows.Next() {
			err := rows.Scan(&u.UserName, &u.Password, &u.First, &u.Last, &u.Role)
			if err != nil {
				panic(err)
			}
		}
	}
	return u
}
