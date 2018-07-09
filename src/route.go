package main

import (
	"database/sql"

	"./controller"
	"./repository"

	"github.com/julienschmidt/httprouter"
)

// SetupIndexRoutes - Routes for /
func SetupIndexRoutes(router *httprouter.Router) {
	indexController := controller.IndexController{}
	router.GET("/", indexController.IndexRoute)
}

// SetupNoteRoutes - Routes for /notes
func SetupNoteRoutes(db *sql.DB, router *httprouter.Router) {
	noteRepo := repository.NoteRepository{DB: db}
	noteContoller := controller.NoteController{NoteRepo: &noteRepo}

	router.GET("/note/:id", noteContoller.GetNote)
	router.PUT("/note/:id", noteContoller.UpdateNote)
	router.DELETE("/note/:id", noteContoller.DeleteNote)
	router.POST("/note", noteContoller.CreateNote)
}
