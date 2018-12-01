package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	dbHOST = "db"
	dbUSER = "mypguser"
	dbPASS = "password"
	dbNAME = "mydb"
)

func connectToDB() *sql.DB {
	connStr := "host=" + dbHOST + " user=" + dbUSER + " dbname=" + dbNAME + " password=" + dbPASS + " sslmode=disable"
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to Database")
	}

	return sqlDB
}

func main() {
	router := httprouter.New()
	db := connectToDB()

	SetupRoutes(db, router)

	fmt.Println("Starting Server at : 8080")
	err := http.ListenAndServe(":8080", router)
	fmt.Println(err.Error())
}
