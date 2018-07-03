package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Note - Holds the Note Data
type Note struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

// NewNote - Creates and returns a new Note
func NewNote(text string) *Note {
	note := Note{
		ID:        1,
		Text:      text,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}

	return &note
}

// Update - Update the text for Note and time
func (n *Note) Update(text string) {
	n.Text = text
	n.UpdatedOn = time.Now()
}

func indexRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome")
}

func createNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	noteText := r.PostForm.Get("note")
	note := NewNote(noteText)
	json, _ := json.Marshal(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", json)
}

func getNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Get Note %s", id)
}

func updateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Update Note %s", id)
}

func deleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Delete Note %s", id)
}

func main() {
	router := httprouter.New()

	router.GET("/", indexRoute)

	router.GET("/note/:id", getNote)
	router.PUT("/note/:id", updateNote)
	router.DELETE("/note/:id", deleteNote)
	router.POST("/note", createNote)

	log.Fatal(http.ListenAndServe(":8080", router))
}
