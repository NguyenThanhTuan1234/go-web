package driver

import (
	"go-web/adapter/controller"
	"go-web/adapter/gateway"
	"go-web/usecase"
	"net/http"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	param := &controller.ParamControllerPreset{
		Param: "session",
	}
	sessionRepo := gateway.NewSessionClient()
	use := usecase.NewLogOutUsecase(param, sessionRepo)
	err := use.CreateLogOutPage(w, r)
	if err != nil {
		panic(err)
	}
}
