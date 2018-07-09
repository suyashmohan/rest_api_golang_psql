package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../repository"
	"./request"

	"github.com/julienschmidt/httprouter"
)

// NoteController - Controller for Note API
type NoteController struct {
	NoteRepo *repository.NoteRepository
}

// CreateNote - Create a new Note in DB
func (nc *NoteController) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	noteReq := request.NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	note := nc.NoteRepo.New(noteReq.Text)
	json, _ := json.Marshal(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", json)
}

// GetNote - Get a note based on ID
func (nc *NoteController) GetNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	note := nc.NoteRepo.Get(id)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}

// UpdateNote - Update a note using ID
func (nc *NoteController) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	noteReq := request.NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	note := nc.NoteRepo.Update(id, noteReq.Text)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}

// DeleteNote - Delete a note using ID
func (nc *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	note := nc.NoteRepo.Delete(id)
	if note != nil {
		json, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", json)
	} else {
		w.WriteHeader(404)
	}
}
