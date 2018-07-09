package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controller"
	"./repository"

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

func setupIndexRoutes(router *httprouter.Router) {
	indexController := controller.IndexController{}
	router.GET("/", indexController.IndexRoute)
}

func setupNoteRoutes(db *sql.DB, router *httprouter.Router) {
	noteRepo := repository.NoteRepository{DB: db}
	noteContoller := controller.NoteController{NoteRepo: &noteRepo}

	router.GET("/note/:id", noteContoller.GetNote)
	router.PUT("/note/:id", noteContoller.UpdateNote)
	router.DELETE("/note/:id", noteContoller.DeleteNote)
	router.POST("/note", noteContoller.CreateNote)
}

func main() {
	router := httprouter.New()
	db := connectToDB()
	if db == nil {
		log.Fatal("Unable to connect to DB")
	}

	setupIndexRoutes(router)
	setupNoteRoutes(db, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
