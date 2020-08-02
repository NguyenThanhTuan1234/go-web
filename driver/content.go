package driver

import (
	"go-web/adapter/gateway"
	"go-web/config"
	"go-web/usecase"
	"net/http"
)

func Content(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	rdbClient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbClient)
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewContentUsecase(postgresRepo, handlerRepo)
	err1 := use.CreateContentPage(w, r)
	if err1 != nil {
		panic(err1)
	}
}
