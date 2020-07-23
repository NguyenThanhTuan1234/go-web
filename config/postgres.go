package config

type postgresConfig struct {
	Port         int    `json: "port"`
	User         string `json: "user"`
	Password     string `json: "password"`
	DatabaseName string `json:"database_name"`
	SSLMode      string `json:"ssl_mode"`
}

func (c *postgresConfig) GetPort() int {
	return c.Port
}

func (c *postgresConfig) GetUser() string {
	return c.User
}

func (c *postgresConfig) GetPassword() string {
	return c.Password
}

func (c *postgresConfig) GetDatabaseName() string {
	return c.DatabaseName
}

func (c *postgresConfig) GetSSLMode() string {
	return c.SSLMode
}
