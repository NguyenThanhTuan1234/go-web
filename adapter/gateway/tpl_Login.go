package gateway

import "net/http"

func (t *templateClient) LogIn(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
		return err
	}
	return nil
}
