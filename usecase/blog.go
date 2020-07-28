package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type BlogUsecase interface {
	service.BlogService
}

type blogUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewBlogUsecase(
	handlerRepo repository.HandlerRepository,
) BlogUsecase {
	return &blogUsecase{
		handlerRepo: handlerRepo,
	}
}
