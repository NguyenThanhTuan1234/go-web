package usecase

import (
	"go-web/entity/repository"
	"go-web/entity/service"
)

type BlogUsecase interface {
	service.BlogService
}

type blogUsecase struct {
	postgresRepo repository.PostgresRepository
	handlerRepo  repository.HandlerRepository
}

func NewBlogUsecase(
	postgresRepo repository.PostgresRepository,
	handlerRepo repository.HandlerRepository,
) BlogUsecase {
	return &blogUsecase{
		postgresRepo: postgresRepo,
		handlerRepo:  handlerRepo,
	}
}
