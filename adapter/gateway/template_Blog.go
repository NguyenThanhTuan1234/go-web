package gateway

import "net/http"

func (t *templateClient) Blog(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		return err
	}
	return nil
}
