package gateway

import (
	"go-web/models"
	"net/http"
)

func (t *templateClient) Index(w http.ResponseWriter, r *http.Request, user models.User) error {
	err := t.tpl.ExecuteTemplate(w, "index.html", user)
	if err != nil {
		return err
	}
	return nil
}
