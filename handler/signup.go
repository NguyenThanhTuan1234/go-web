package handler

import (
	"go-web/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// var u models.User
	// process form submission
	if r.Method == http.MethodPost {

		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		role := r.FormValue("role")

		// username taken?
		// if _, ok := dbUsers[un]; ok {
		// 	http.Error(w, "Username already taken", http.StatusForbidden)
		// 	return
		// }
		// username taken ?
		rows, err := db.Query("SELECT username FROM test1 WHERE username = $1", un)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		var u models.User
		for rows.Next() {
			err := rows.Scan(&u.UserName)
			if err != nil {
				panic(err)
			}
		}
		if u.UserName != "" {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = models.Session{u.ID, time.Now()}

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// u := models.NewUser(un, bs, f, l, role)
		// dbUsers[un] = u

		_, err = db.Exec("INSERT INTO test1 (username, password, first, last, role) VALUES ($1, $2, $3, $4, $5)", un, string(bs), f, l, role)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.html", nil)
}
