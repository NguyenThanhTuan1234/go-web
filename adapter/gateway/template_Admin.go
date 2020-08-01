package gateway

import (
	"go-web/models"
	"net/http"
)

func (t *templateClient) Admin(w http.ResponseWriter, r *http.Request, user models.User) error {
	err := t.tpl.ExecuteTemplate(w, "admin.html", user)
	if err != nil {
		return err
	}
	return nil
}
