package driver

import (
	"go-web/adapter/controller"
	"go-web/adapter/gateway"
	"go-web/config"
	"go-web/usecase"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	formIn1 := &controller.FormControllerPreset{
		Value: "username",
	}
	formIn2 := &controller.FormControllerPreset{
		Value: "password",
	}
	rdbclient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbclient)
	if err != nil {
		panic(err)
	}
	bcryptRepo := gateway.NewBcrypt()
	sessionRepo := gateway.NewSessionClient()
	sessionName := &controller.ParamControllerPreset{
		Param: "session",
	}
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewLoginUsecase(formIn1, formIn2, postgresRepo, bcryptRepo, sessionName, sessionRepo, handlerRepo)
	err1 := use.CreateLoginPage(w, r)
	if err1 != nil {
		panic(err)
	}
}
