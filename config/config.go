package config

import (
	"go-web/adapter/gateway/client"
)

type Config interface {
	GetPostgres() client.PostgresConfig
}

var conf Config

func GetConfig() Config {
	return conf
}

type config struct {
	Postgres *postgresConfig `json:"postgres"`
}

func (c *config) GetPostgres() client.PostgresConfig {
	return c.Postgres
}
