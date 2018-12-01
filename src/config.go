package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// Config -
type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"db"`
	}

	App struct {
		Port string `yaml:"port"`
	}
}

// Read config file
func (c *Config) Read(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading config file")
		fmt.Println(err)
		c.loadDefault()
	}
	yaml.Unmarshal([]byte(data), c)
}

// DBString return Connection String for DB
func (c *Config) DBString() string {
	return "host=" + c.DB.Host + " port=" + c.DB.Port + " user=" + c.DB.User + " dbname=" + c.DB.Name + " password=" + c.DB.Password + " sslmode=disable"
}

func (c *Config) loadDefault() {
	c.DB.Host = "db"
	c.DB.User = "mypguser"
	c.DB.Password = "password"
	c.DB.Name = "mydb"
	c.App.Port = "8080"
}
