package usecase

import "net/http"

type ValueInput interface {
	GetFormValue(http.ResponseWriter, *http.Request, string) string
}
