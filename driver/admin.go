package driver

import (
	"go-web/adapter/controller"
	"go-web/adapter/gateway"
	"go-web/config"
	"go-web/usecase"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	formIn1 := &controller.FormControllerPreset{
		Value: "content",
	}
	formIn2 := &controller.FormControllerPreset{
		Value: "title",
	}
	sessionRepo := gateway.NewSessionClient()
	handlerRepo := gateway.NewTemplateRepository()
	rdbclient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbclient)
	use := usecase.NewAdminUsecase(formIn1, formIn2, sessionRepo, postgresRepo, handlerRepo)
	err1 := use.CreateAdminPage(w, r)
	if err1 != nil {
		panic(err)
	}
}
