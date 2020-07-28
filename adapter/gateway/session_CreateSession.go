package gateway

import (
	"fmt"
	"go-web/models"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const sessionLength int = 600

var session *models.Session
var dbSessions = map[string]*models.Session{}

func (c *sessionRepository) CreateSession(w http.ResponseWriter, r *http.Request, name string, un string) error {

	sID, err := uuid.NewV4()
	fmt.Println(sID.String())
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   name,
		Value:  sID.String(),
		MaxAge: sessionLength,
	}
	http.SetCookie(w, cookie)
	dbSessions[cookie.Value] = &models.Session{un, time.Now()}
	return nil
}
