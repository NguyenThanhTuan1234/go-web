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
			return err1
		}
		sessionname, err2 := l.param.ReadParam()
		if err2 != nil {
			return err2
		}
		err3 := l.sessionRepo.CreateSession(w, r, sessionname)
		if err3 != nil {
			return err3
		}
		err4 := l.handlerRepo.LogIn(w, r)
		if err4 != nil {
			return err4
		}
	}
	return nil
}
