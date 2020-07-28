package service

import "net/http"

type BlogService interface {
	CreateBlogPage(w http.ResponseWriter, r *http.Request) error
}