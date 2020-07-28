package gateway

import "net/http"

func (t *templateClient) SignUp(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		return err
	}
	return nil
}
