package gateway

import (
	"go-web/models"
	"net/http"
)

func (c *sessionRepository) GetSessionInfo(w http.ResponseWriter, r *http.Request) (*models.Session, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}
	return dbSessions[cookie.Value], nil
}
