package usecase

import (
	"fmt"
	"net/http"
)

func (l *loginUsecase) CreateLoginPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err2 := l.param.ReadParam()
	if err2 != nil {
		return err2
	}
	fmt.Println(sessionname)
	if l.sessionRepo.CheckSessionIfExist(w, r, sessionname) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
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
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return nil
		}
		err1 := l.bcryptRepo.ComparePassword(user.GetPassWord(), password)
		if err1 != nil {
			fmt.Println("test")
			http.Error(w, "Username or password do not match", http.StatusForbidden)
			return nil
		}
		sessionname, err2 := l.param.ReadParam()
		if err2 != nil {
			return err2
		}
		err3 := l.sessionRepo.CreateSession(w, r, sessionname, username)
		if err3 != nil {
			return err3
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err4 := l.handlerRepo.LogIn(w, r)
	if err4 != nil {
		return err4
	}
	return nil
}
