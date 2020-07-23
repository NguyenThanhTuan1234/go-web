package client

import (
	"database/sql"
	"go-web/models"
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

type RDBClient interface {
	GetUser(string) (*models.User, error)
	CreateUser(string, string, string, string, []byte) error
}

type rdbClient struct {
	db *sql.DB
}

func NewRDBClient(config PostgresConfig) (RDBClient, error) {
	specific := NewPostgresSpecific(config)
	db, err := sql.Open(specific.GetDriver(), specific.GetSources())
	if err != nil {
		return nil, err
	}
	return &rdbClient{
		db: db,
	}, nil
}
