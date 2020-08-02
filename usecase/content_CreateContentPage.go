package usecase

import (
	"go-web/models"
	"net/http"
)

func (u *contentUsecase) CreateContentPage(w http.ResponseWriter, r *http.Request) error {
	content, err := u.postgresRepo.GetContent()
	if err != nil {
		return err
	}
	err1 := u.handlerRepo.Content(w, r, models.NewContent(content.Title, content.Content, content.Uid))
	if err1 != nil {
		return err1
	}
	return nil
}
