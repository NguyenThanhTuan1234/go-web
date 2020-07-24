package gateway

import "net/http"

func (t *templateClient) Index(w http.ResponseWriter, r *http.Request) {
	t.tpl.ExecuteTemplate(w, "index.html", nil)
}
