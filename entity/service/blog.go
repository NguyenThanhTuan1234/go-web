package service

import "net/http"

type BlogService interface {
	CreateBlogPage(http.ResponseWriter, *http.Request) error
}
