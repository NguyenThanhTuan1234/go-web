package controller

import "net/http"

type ValueControllerPreset struct {
	_    struct{}
	Name string
}

func (c *ValueControllerPreset) GetFormValue(w http.ResponseWriter, r *http.Request, x string) string {
	c.Name = r.FormValue(x)
	return c.Name
}
