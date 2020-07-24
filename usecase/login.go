package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type LoginUsecase interface {
	service.LoginService
}

type loginUsecase struct {
	postgresRepo repository.PostgresRepository
	handlerRepo  repository.HandlerRepository
	valueIn      ValueInput
}
