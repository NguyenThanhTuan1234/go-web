package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type LogOutUsecase interface {
	service.LogOutService
}

type logoutUseCase struct {
	param       ParamInput
	sessionRepo repository.SessionRepository
}

func NewLogOutUsecase(
	param ParamInput,
	sessionRepo repository.SessionRepository,
) LogOutUsecase {
	return &logoutUseCase{
		param:       param,
		sessionRepo: sessionRepo,
	}
}
