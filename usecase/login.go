package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type LoginUsecase interface {
	service.LoginService
}

type loginUsecase struct {
	formIn1      FormInput
	formIn2      FormInput
	postgresRepo repository.PostgresRepository
	bcryptRepo   repository.BcryptRepository
	sessionRepo  repository.SessionRepository
	handlerRepo  repository.HandlerRepository
}

func NewLoginUsecase(
	formIn1 FormInput,
	formIn2 FormInput,
	postgresRepo repository.PostgresRepository,
	bcryptRepo repository.BcryptRepository,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) LoginUsecase {
	return &loginUsecase{
		formIn1:      formIn1,
		formIn2:      formIn2,
		postgresRepo: postgresRepo,
		bcryptRepo:   bcryptRepo,
		sessionRepo:  sessionRepo,
		handlerRepo:  handlerRepo,
	}
}
