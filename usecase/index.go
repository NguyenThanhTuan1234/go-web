package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type IndexUsecase interface {
	service.IndexService
}

type indexUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewIndexUsecase(
	handlerRepo repository.HandlerRepository,
) IndexUsecase {
	return &indexUsecase{
		handlerRepo: handlerRepo,
	}
}
