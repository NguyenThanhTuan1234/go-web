package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		err = fmt.Errorf("confgen.main() read config file error, %w", err)
		panic(err)
	}
	src := fmt.Sprintf(`
	package config

import "encoding/json"

func init() {
	conf = new(config)
	b := []byte(configData)
	if err := json.Unmarshal(b, conf); err != nil {
		panic(err)
	}
}

const configData = %s
`, "`"+string(bytes)+"`")
	bin, err := format.Source([]byte(src))
	if err != nil {
		log.Fatalf("go src format: %s", err)
	} else if err = ioutil.WriteFile("config_init.go", bin, 0644); err != nil {
		log.Fatalf("Writing output: %s", err)
	}

}
