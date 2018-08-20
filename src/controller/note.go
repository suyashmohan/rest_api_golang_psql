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
func (nc *NoteController) CreateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	noteReq := request.NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	if len(strings.TrimSpace(noteReq.Text)) == 0 {
		response.BadRequest("Note is empty", w)
	} else {
		userID := ps.ByName("userid")
		note := nc.NoteRepo.New(noteReq.Text, userID)
		response.Success(note, w)
	}
}

// GetNote - Get a note based on ID
func (nc *NoteController) GetNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	userID := ps.ByName("userid")
	note := nc.NoteRepo.Get(id, userID)
	if note != nil {
		response.Success(note, w)
	} else {
		response.NotFound("Note not found", w)
	}
}

// UpdateNote - Update a note using ID
func (nc *NoteController) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	userID := ps.ByName("userid")

	noteReq := request.NewNoteRequest{}
	json.NewDecoder(r.Body).Decode(&noteReq)

	note := nc.NoteRepo.Update(id, userID, noteReq.Text)
	if note != nil {
		response.Success(note, w)
	} else {
		response.NotFound("Note not found", w)
	}
}

// DeleteNote - Delete a note using ID
func (nc *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	userID := ps.ByName("userid")
	note := nc.NoteRepo.Delete(id, userID)
	if note != nil {
		response.Success(note, w)
	} else {
		response.NotFound("Note not found", w)
	}
}

// AuthError - Send Error for Auth
func (nc *NoteController) AuthError(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response.Unauthorized("The request requires authorization", w)
}
