package usecase

import "net/http"

func (u *blogUsecase) CreateBlogPage(w http.ResponseWriter, r *http.Request) error {
	err := u.handlerRepo.Blog(w, r)
	if err != nil {
		return err
	}
	return nil
}
