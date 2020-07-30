package usecase

import (
	"go-web/models"
	"net/http"
)

func (i *indexUsecase) CreateIndexPage(w http.ResponseWriter, r *http.Request) error {
	err := i.handlerRepo.Index(w, r, models.NewUser(0, "", nil, "", "", ""))
	if err != nil {
		return err
	}
	return nil
}
