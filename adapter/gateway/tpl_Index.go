package gateway

import "net/http"

func (t *templateClient) Index(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return err
	}
	return nil
}
