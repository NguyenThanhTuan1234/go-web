package usecase

import "net/http"

func (i *indexUsecase) CreateIndexPage(w http.ResponseWriter, r *http.Request) error {
	err := i.handlerRepo.Index(w, r)
	if err != nil {
		return err
	}
	return nil
}
