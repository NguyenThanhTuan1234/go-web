package gateway

import "go-web/entity/repository"

type PostgresConfig interface {
	GetUser() string
	GetPassword() string
	GetDatabaseName() string
	GetSSLMode() string
}

type PostgresRepository interface {
	repository.PostgresRepository
}

type postgresRepository struct {
	client *rdbClient
}

func NewPostgresRepository(client *rdbClient) PostgresRepository {
	return &postgresRepository{
		client: client,
	}
}
