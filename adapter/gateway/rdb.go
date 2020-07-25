package gateway

import (
	"database/sql"
	"go-web/entity/repository"
)

type RDBSpecific interface {
	GetDriver() string
	GetSources() string
}

type rdbSpecific struct {
	driver  string
	sources string
}

func (s *rdbSpecific) GetDriver() string  { return s.driver }
func (s *rdbSpecific) GetSources() string { return s.sources }

type RDBRepository interface {
	repository.PostgresRepository
}

type rdbClient struct {
	db *sql.DB
}

func NewRDBRepository(config PostgresConfig) (RDBRepository, error) {
	specific := NewPostgresSpecific(config)
	db, err := sql.Open(specific.GetDriver(), specific.GetSources())
	if err != nil {
		return nil, err
	}
	return &rdbClient{
		db: db,
	}, nil
}
