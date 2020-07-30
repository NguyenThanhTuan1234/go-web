package repository

import (
	"net/http"
)

type SessionRepository interface {
	CreateSession(http.ResponseWriter, *http.Request, string, int) error
	CheckSessionIfExist(http.ResponseWriter, *http.Request, string) bool
	DeleteSession(http.ResponseWriter, *http.Request, string) error
}
