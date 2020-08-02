package usecase

import (
	"go-web/models"
	"net/http"
)

func (u *adminUsecase) CreateAdminPage(w http.ResponseWriter, r *http.Request) error {
	session, err := u.sessionRepo.GetSessionInfo(w, r)
	if err != nil {
		return err
	}
	if r.Method == http.MethodPost {
		content, err := u.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		title, err := u.formIn2.GetFormValue(w, r)
		if err != nil {
			return nil
		}
		err1 := u.postgresRepo.CreateContent(content, title, session.Id)
		if err1 != nil {
			return err
		}
	}
	err1 := u.handlerRepo.Admin(w, r, models.NewUser(0, "", nil, "", "", ""))
	if err1 != nil {
		return err1
	}
	return nil
}
