package handler

import (
	"go-web/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		// is there a username ?
		// u, ok := dbUsers[un]
		// if !ok {
		// 	http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		// 	return
		// }
		var u models.User
		rows, err := db.Query("SELECT username, password, first, last, role FROM test1 WHERE username = $1", un)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		}
		for rows.Next() {
			err := rows.Scan(&u.UserName, &u.Password, &u.First, &u.Last, &u.Role)
			if err != nil {
				panic(err)
			}
		}
		// does the entered password match the stored password
		err1 := bcrypt.CompareHashAndPassword(u.GetPassWord(), []byte(p))
		if err1 != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.html", nil)
}
