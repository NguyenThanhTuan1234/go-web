package gateway

import "net/http"

func (c *sessionRepository) DeleteSession(w http.ResponseWriter, r *http.Request, name string) error {
	cookie, err := r.Cookie(name)
	if err != nil {
		return err
	}
	delete(dbSessions, cookie.Value)

	cookie = &http.Cookie{
		Name:   name,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	return nil
}
