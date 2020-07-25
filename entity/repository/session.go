package repository

import (
	"net/http"
)

type SessionRepository interface {
	CreateSession(http.ResponseWriter, *http.Request, string) error
}
