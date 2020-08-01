package usecase

import (
	"go-web/models"
	"net/http"
)

func (u *adminUsecase) CreateAdminPage(w http.ResponseWriter, r *http.Request) error {
	err := u.handlerRepo.Admin(w, r, models.NewUser(0, "", nil, "", "", ""))
	if err != nil {
		return err
	}
	return nil
}
