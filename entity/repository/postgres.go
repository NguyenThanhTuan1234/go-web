package repository

import "go-web/models"

type PostgresRepository interface {
	GetUser(string) (*models.User, error)
	CreateUser(string, string, string, string, string) error
}
