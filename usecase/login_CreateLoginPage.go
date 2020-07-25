package usecase

import "net/http"

func (l *loginUsecase) CreateLoginPage(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		username, err := l.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		password, err := l.formIn2.GetFormValue(w, r)
		if err != nil {
			return err
		}
		user, err := l.postgresRepo.GetUser(username)
		if err != nil {
			return err
		}
		err1 := l.bcryptRepo.ComparePassword(user.GetPassWord(), password)
		if err1 != nil {
			return err
		}
		err2 := l.sessionRepo.CreateSession(w, r, "session")
		if err != nil {
			return err2
		}
		err3 := l.handlerRepo.LogIn(w, r)
		if err != nil {
			return err3
		}
	}
	return nil
}
