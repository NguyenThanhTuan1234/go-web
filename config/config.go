package config

import (
	"go-web/adapter/gateway"
)

type Config interface {
	GetPostgres() gateway.PostgresConfig
}

var conf Config

func GetConfig() Config {
	return conf
}

type config struct {
	Postgres *postgresConfig `json:"postgres"`
}

func (c *config) GetPostgres() gateway.PostgresConfig {
	return c.Postgres
}
