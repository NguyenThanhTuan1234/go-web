package gateway

import (
	"go-web/models"
	"net/http"
)

func (t *templateClient) Blog(w http.ResponseWriter, r *http.Request, content models.Content) error {
	err := t.tpl.ExecuteTemplate(w, "blog.html", content)
	if err != nil {
		return err
	}
	return nil
}
