package usecase

import "net/http"

func (s *signupUsecase) CreateSignUpPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err := s.param.ReadParam()
	if err != nil {
		return err
	}
	if s.sessionRepo.CheckSessionIfExist(w, r, sessionname) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
	if r.Method == http.MethodPost {
		username, err := s.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		password, err := s.formIn2.GetFormValue(w, r)
		if err != nil {
			return err
		}
		first, err := s.formIn3.GetFormValue(w, r)
		if err != nil {
			return err
		}
		last, err := s.formIn4.GetFormValue(w, r)
		if err != nil {
			return err
		}
		role, err := s.formIn5.GetFormValue(w, r)
		if err != nil {
			return err
		}
		user, err := s.postgresRepo.GetUser(username)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return nil
		}
		if user.GetUserName() != "" {
			http.Error(w, "Username already taken", http.StatusForbidden)
		}
		sessionname, err := s.param.ReadParam()
		if err != nil {
			return err
		}
		err2 := s.sessionRepo.CreateSession(w, r, sessionname, user.GetId())
		if err != nil {
			return err2
		}
		hashPassword, err := s.bcryptRepo.GenerateHashPassword(password)
		if err != nil {
			return err
		}
		err3 := s.postgresRepo.CreateUser(username, hashPassword, first, last, role)
		if err3 != nil {
			return err3
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err4 := s.handlerRepo.SignUp(w, r)
	if err4 != nil {
		return err4
	}
	return nil
}
