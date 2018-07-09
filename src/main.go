package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	dbUSER = "mypguser"
	dbPASS = "password"
	dbNAME = "mydb"
)

func connectToDB() *sql.DB {
	connStr := "user=" + dbUSER + " dbname=" + dbNAME + " password=" + dbPASS + " sslmode=disable"
	sqlDB, _ := sql.Open("postgres", connStr)
	return sqlDB
}

func main() {
	router := httprouter.New()
	db := connectToDB()
	if db == nil {
		log.Fatal("Unable to connect to DB")
	}

	SetupIndexRoutes(router)
	SetupNoteRoutes(db, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
