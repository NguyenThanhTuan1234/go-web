package gateway

import "go-web/entity/repository"

type sessionRepository struct{}

type SessionRepository interface {
	repository.SessionRepository
}

func NewSessionClient() SessionRepository {
	return &sessionRepository{}
}
