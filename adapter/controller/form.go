package controller

import "net/http"

type FormControllerPreset struct {
	_     struct{}
	Value string
}

func (c *FormControllerPreset) GetFormValue(w http.ResponseWriter, r *http.Request) (string, error) {
	return r.FormValue(c.Value), nil
}
