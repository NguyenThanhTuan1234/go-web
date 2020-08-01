package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type AdminUsecase interface {
	service.AdminService
}

type adminUsecase struct {
	formIn1      FormInput
	formIn2      FormInput
	sessionRepo  repository.SessionRepository
	postgresRepo repository.PostgresRepository
	handlerRepo  repository.HandlerRepository
}

func NewAdminUsecase(
	formIn1 FormInput,
	formIn2 FormInput,
	sessionRepo repository.SessionRepository,
	postgresRepo repository.PostgresRepository,
	handlerRepo repository.HandlerRepository,
) AdminUsecase {
	return &adminUsecase{
		formIn1:      formIn1,
		formIn2:      formIn2,
		sessionRepo:  sessionRepo,
		postgresRepo: postgresRepo,
		handlerRepo:  handlerRepo,
	}
}
