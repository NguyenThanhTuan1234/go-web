package gateway

import "fmt"

type PostgresConfig interface {
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabaseName() string
	GetSSLMode() string
}

func NewPostgresSpecific(c PostgresConfig) RDBSpecific {
	return &rdbSpecific{
		driver: "postgres",
		sources: fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.GetPort(), c.GetUser(), c.GetPassword(), c.GetDatabaseName(), c.GetSSLMode()),
	}
}
