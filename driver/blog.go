package driver

import (
	"go-web/adapter/gateway"
	"go-web/usecase"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewBlogUsecase(handlerRepo)
	err := use.CreateBlogPage(w, r)
	if err != nil {
		panic(err)
	}
}