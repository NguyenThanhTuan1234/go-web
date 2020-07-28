package gateway

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type RDBSpecific interface {
	GetDriver() string
	GetSource() string
}

type rdbSpecific struct {
	driver string
	source string
}

func (s *rdbSpecific) GetDriver() string { return s.driver }
func (s *rdbSpecific) GetSource() string { return s.source }

func NewPostgresSpecific(c PostgresConfig) RDBSpecific {
	return &rdbSpecific{
		driver: "postgres",
		source: fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", c.GetUser(), c.GetPassword(), c.GetDatabaseName(), c.GetSSLMode()),
	}
}

type rdbClient struct {
	db *sql.DB
}

func NewRDBClient(config PostgresConfig) (*rdbClient, error) {
	specific := NewPostgresSpecific(config)
	fmt.Println(specific.GetDriver())
	fmt.Println(specific.GetSource())
	db, err := sql.Open(specific.GetDriver(), specific.GetSource())
	if err != nil {
		return nil, err
	}
	fmt.Println("You connected to your database.")
	return &rdbClient{
		db: db,
	}, nil
}
