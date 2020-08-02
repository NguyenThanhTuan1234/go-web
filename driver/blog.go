package driver

import (
	"go-web/adapter/gateway"
	"go-web/config"
	"go-web/usecase"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	handlerRepo := gateway.NewTemplateRepository()
	rdbclient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbclient)
	use := usecase.NewBlogUsecase(postgresRepo, handlerRepo)
	err1 := use.CreateBlogPage(w, r)
	if err1 != nil {
		panic(err1)
	}
}
