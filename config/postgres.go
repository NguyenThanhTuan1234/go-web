package config

type postgresConfig struct {
	User         string `json: "user"`
	Password     string `json: "password"`
	DatabaseName string `json:"databasename"`
	SSLMode      string `json:"sslmode"`
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
