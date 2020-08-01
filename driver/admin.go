package driver

import (
	"go-web/adapter/gateway"
	"go-web/usecase"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewAdminUsecase(handlerRepo)
	err := use.CreateAdminPage(w, r)
	if err != nil {
		panic(err)
	}
}
