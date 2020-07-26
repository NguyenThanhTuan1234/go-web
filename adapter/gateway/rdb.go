package gateway

import (
	"database/sql"
	"fmt"
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
	fmt.Println(specific.GetDriver())
	fmt.Println(specific.GetSources())
	db, err := sql.Open(specific.GetDriver(), specific.GetSources())
	if err != nil {
		return nil, err
	}
	fmt.Println("You connected to your database.")
	return &rdbClient{
		db: db,
	}, nil
}

func NewPostgresSpecific(c PostgresConfig) RDBSpecific {
	return &rdbSpecific{
		driver: "postgres",
		sources: fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.GetPort(), c.GetUser(), c.GetPassword(), c.GetDatabaseName(), c.GetSSLMode()),
	}
}
