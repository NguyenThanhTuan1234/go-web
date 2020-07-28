package usecase

import "net/http"

func (l *logoutUseCase) CreateLogOutPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err := l.param.ReadParam()
	if err != nil {
		return err
	}
	if l.sessionRepo.CheckSessionIfExist(w, r, sessionname) == false {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
	err1 := l.sessionRepo.DeleteSession(w, r, sessionname)
	if err1 != nil {
		return err1
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}
