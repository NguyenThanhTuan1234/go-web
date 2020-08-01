package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type AdminUsecase interface {
	service.AdminService
}

type adminUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewAdminUsecase(handlerRepo repository.HandlerRepository) AdminUsecase {
	return &adminUsecase{
		handlerRepo: handlerRepo,
	}
}
