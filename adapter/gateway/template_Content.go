package gateway

import (
	"go-web/models"
	"net/http"
)

func (t *templateClient) Content(w http.ResponseWriter, r *http.Request, content models.Content) error {
	err := t.tpl.ExecuteTemplate(w, "blog_content.html", content)
	if err != nil {
		return err
	}
	return nil
}
