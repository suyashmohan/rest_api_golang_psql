package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	dbUSER = "mypguser"
	dbPASS = "password"
	dbNAME = "mydb"
)

var db *sql.DB

// Note - Holds the Note Data
type Note struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

// NewNoteRequest - Request Body for New Note
type NewNoteRequest struct {
	Text string `json:"text"`
}

// NoteRepository - Repository for Database Layer of Notes
type NoteRepository struct{}

func connectToDB() *sql.DB {
	connStr := "user=" + dbUSER + " dbname=" + dbNAME + " password=" + dbPASS + " sslmode=disable"
	sqlDB, _ := sql.Open("postgres", connStr)
	return sqlDB
}

// New - Create new Note in DB
func (nr *NoteRepository) New(text string) *Note {
	dbNote := Note{}
	row := db.QueryRow("INSERT INTO notes(text, createdOn, updatedOn) VALUES($1::TEXT, now()::TIMESTAMP, now()::TIMESTAMP) RETURNING *", text)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err != nil {
		log.Fatal(err)
	}

	return &dbNote
}

// Get - Return a Note based on ID
func (nr *NoteRepository) Get(id string) *Note {
	dbNote := Note{}
	row := db.QueryRow("SELECT * FROM notes WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err == nil {
		return &dbNote
	}

	return nil
}

// Update - Update the text for Note
func (nr *NoteRepository) Update(id, text string) *Note {
	dbNote := Note{}
	row := db.QueryRow("UPDATE notes SET text=$2, updatedon=now()::TIMESTAMP WHERE id=$1 RETURNING *", id, text)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)
	if err == nil {
		return &dbNote
	}

	return nil
}

// Delete - Delete a record from DB
func (nr *NoteRepository) Delete(id string) *Note {
	dbNote := Note{}
	row := db.QueryRow("DELETE FROM notes WHERE id=$1 RETURNING *", id)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)
	if err == nil {
		return &dbNote
	}

	return nil
}

func indexRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello! World")
}

func createNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	noteReq := NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	noteRepo := NoteRepository{}
	note := noteRepo.New(noteReq.Text)
	json, _ := json.Marshal(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", json)
}

func getNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	noteRepo := NoteRepository{}
	note := noteRepo.Get(id)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}

func updateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	noteReq := NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	noteRepo := NoteRepository{}
	note := noteRepo.Update(id, noteReq.Text)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}

func deleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	noteRepo := NoteRepository{}
	note := noteRepo.Delete(id)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}

func main() {
	router := httprouter.New()
	db = connectToDB()
	if db == nil {
		log.Fatal("Unable to connect to DB")
	}

	router.GET("/", indexRoute)

	router.GET("/note/:id", getNote)
	router.PUT("/note/:id", updateNote)
	router.DELETE("/note/:id", deleteNote)
	router.POST("/note", createNote)

	log.Fatal(http.ListenAndServe(":8080", router))
}
