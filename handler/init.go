package handler

import (
	"database/sql"
	"fmt"
	"go-web/models"
	"html/template"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "login-golang"
	SSL         = "disable"
)

var db *sql.DB
var tpl *template.Template

// var dbUsers = map[string]models.User{}       // username, user
var dbSessions = map[string]models.Session{} // session ID, username
var dbSessionsCleaned time.Time

const sessionLength int = 600

func init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, DB_NAME, SSL)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
	tpl = template.Must(template.ParseGlob("public/*"))
}
