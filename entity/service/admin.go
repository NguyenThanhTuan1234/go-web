package service

import "net/http"

type AdminService interface {
	CreateAdminPage(http.ResponseWriter, *http.Request) error
}
