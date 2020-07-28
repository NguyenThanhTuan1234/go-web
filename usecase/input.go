package usecase

import "net/http"

type FormInput interface {
	GetFormValue(http.ResponseWriter, *http.Request) (string, error)
}

type ParamInput interface {
	ReadParam() (string, error)
}
