package handler

import (
	"go-web/models"
	"html/template"
	"time"
)

var tpl *template.Template
var dbUsers = map[string]models.User{}       // user ID, user
var dbSessions = map[string]models.Session{} // session ID, user ID
var dbSessionsCleaned time.Time

const sessionLength int = 600

func init() {
	tpl = template.Must(template.ParseGlob("public/*"))
}
