package config

import "encoding/json"

func init() {
	conf = new(config)
	b := []byte(configData)
	if err := json.Unmarshal(b, conf); err != nil {
		panic(err)
	}
}

const configData = `{
    "postgres": {
        "user": "postgres",
        "password": "postgres",
        "databasename": "login-golang",
        "sslmode": "disable"
    }
}`
