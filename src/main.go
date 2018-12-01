package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func connectToDB(c Config) *sql.DB {
	connStr := c.DBString()
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to Database")
	}

	return sqlDB
}

func main() {
	config := Config{}
	config.Read("config.yml")

	router := httprouter.New()
	db := connectToDB(config)

	SetupRoutes(db, router)

	fmt.Println("Starting Server at : " + config.App.Port)
	err := http.ListenAndServe(":"+config.App.Port, router)
	fmt.Println(err.Error())
}
