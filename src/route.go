package main

import (
	"database/sql"

	"./controller"
	"./repository"
	"./service"

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

	authReq := func(h httprouter.Handle) httprouter.Handle {
		return service.Auth(h, noteContoller.AuthError)
	}

	router.GET("/note/:id", authReq(noteContoller.GetNote))
	router.PUT("/note/:id", authReq(noteContoller.UpdateNote))
	router.DELETE("/note/:id", authReq(noteContoller.DeleteNote))
	router.POST("/note", authReq(noteContoller.CreateNote))
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
