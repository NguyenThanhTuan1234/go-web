package usecase

import (
	"fmt"
	"go-web/models"
	"net/http"
)

func (l *loginUsecase) CreateLoginPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err := l.param.ReadParam()
	if err != nil {
		return err
	}
	if l.sessionRepo.CheckSessionIfExist(w, r, sessionname) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
	var user *models.User
	var err5 error
	if r.Method == http.MethodPost {
		username, err := l.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		password, err := l.formIn2.GetFormValue(w, r)
		if err != nil {
			return err
		}
		user, err5 = l.postgresRepo.GetUser(username)
		if err5 != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return nil
		}
		err1 := l.bcryptRepo.ComparePassword(user.GetPassWord(), password)
		if err1 != nil {
			fmt.Println("test")
			http.Error(w, "Username or password do not match", http.StatusForbidden)
			return nil
		}
		sessionname, err := l.param.ReadParam()
		if err != nil {
			return err
		}
		err3 := l.sessionRepo.CreateSession(w, r, sessionname, user.GetId())
		if err3 != nil {
			return err3
		}

		fmt.Println(user.UserName)
		l.handlerRepo.Index(w, r, models.NewUser(user.ID, user.UserName, user.Password, user.First, user.Last, user.Role))
	}
	if user == nil {
		err4 := l.handlerRepo.LogIn(w, r)
		if err4 != nil {
			return err4
		}
	}
	return nil
}
