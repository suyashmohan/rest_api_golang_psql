package main

import (
	"database/sql"

	"./controller"
	"./repository"

	"github.com/julienschmidt/httprouter"
)

// SetupIndexRoutes - Routes for /
func setupIndexRoutes(router *httprouter.Router) {
	indexController := controller.IndexController{}
	router.GET("/", indexController.IndexRoute)
}

// SetupNoteRoutes - Routes for /notes
func setupNoteRoutes(db *sql.DB, router *httprouter.Router) {
	noteRepo := repository.NoteRepository{DB: db}
	noteContoller := controller.NoteController{NoteRepo: &noteRepo}

	router.GET("/note/:id", noteContoller.GetNote)
	router.PUT("/note/:id", noteContoller.UpdateNote)
	router.DELETE("/note/:id", noteContoller.DeleteNote)
	router.POST("/note", noteContoller.CreateNote)
}

// SetupUserRoutes - Routes for /user
func setupUserRoutes(db *sql.DB, router *httprouter.Router) {
	userRepo := repository.UserRepository{DB: db}
	userController := controller.UserController{UserRepo: &userRepo}

	router.POST("/user", userController.CreateUser)
	router.POST("/user/verify", userController.VerifyUser)
}

// SetupRoutes - Set up all the Routes
func SetupRoutes(db *sql.DB, router *httprouter.Router) {
	setupIndexRoutes(router)
	setupNoteRoutes(db, router)
	setupUserRoutes(db, router)
}
