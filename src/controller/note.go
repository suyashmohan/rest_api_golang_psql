package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"../repository"
	"./request"
	"./response"

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

	if len(strings.TrimSpace(noteReq.Text)) == 0 {
		response.BadRequest("Note is empty", w)
	} else {
		note := nc.NoteRepo.New(noteReq.Text)
		response.Success(note, w)
	}
}

// GetNote - Get a note based on ID
func (nc *NoteController) GetNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	note := nc.NoteRepo.Get(id)
	if note != nil {
		response.Success(note, w)
	} else {
		response.BadRequest("Note not found", w)
	}
}

// UpdateNote - Update a note using ID
func (nc *NoteController) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	noteReq := request.NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	note := nc.NoteRepo.Update(id, noteReq.Text)
	if note != nil {
		response.Success(note, w)
	} else {
		response.BadRequest("Note not found", w)
	}
}

// DeleteNote - Delete a note using ID
func (nc *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	note := nc.NoteRepo.Delete(id)
	if note != nil {
		response.Success(note, w)
	} else {
		response.BadRequest("Note not found", w)
	}
}
