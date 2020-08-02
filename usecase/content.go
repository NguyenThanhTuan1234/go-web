package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type ContentUsecase interface {
	service.ContentService
}

type contentUsecase struct {
	postgresRepo repository.PostgresRepository
	handlerRepo  repository.HandlerRepository
}

func NewContentUsecase(
	postgresRepo repository.PostgresRepository,
	handlerRepo repository.HandlerRepository,
) ContentUsecase {
	return &contentUsecase{
		postgresRepo: postgresRepo,
		handlerRepo:  handlerRepo,
	}
}
