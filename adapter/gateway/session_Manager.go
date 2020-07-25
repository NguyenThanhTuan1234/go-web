package gateway

import (
	"go-web/entity/repository"
	"go-web/models"
	"net/http"

	uuid "github.com/nu7hatch/gouuid"
)

const sessionLength int = 600

var dbSessions = map[string]*models.Session{}

type sessionRepository struct{}

type SessionRepository interface {
	repository.SessionRepository
}

func NewSessionClient() SessionRepository {
	return &sessionRepository{}
}

var session *models.Session

func (c *sessionRepository) CreateSession(w http.ResponseWriter, r *http.Request, name string) error {
	sID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   name,
		Value:  sID.String(),
		MaxAge: sessionLength,
	}
	http.SetCookie(w, cookie)
	session = dbSessions[cookie.Value]
	return nil
}
