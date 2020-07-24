package gateway

import "net/http"

func (t *templateClient) LogIn(w http.ResponseWriter, r *http.Request) {
	t.tpl.ExecuteTemplate(w, "login", nil)
}
